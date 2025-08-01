package clients

import (
	"context"
	"encoding/json"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/upbound/provider-aws/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/upbound/provider-aws/apis/namespaced/v1beta1"
)

const (
	errNoProviderConfig  = "no providerConfigRef provided"
	errGetProviderConfig = "cannot get referenced ProviderConfig"
	errTrackUsage        = "cannot track ProviderConfig usage"
)

func legacyToModernProviderConfigSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ClusterProviderConfig, error) {
	// TODO(erhan): this is hacky and potentially lossy, generate or manually implement
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ClusterProviderConfig
	err = json.Unmarshal(data, &mSpec)
	mSpec.TypeMeta.Kind = namespacedv1beta1.ClusterProviderConfigKind
	mSpec.TypeMeta.APIVersion = namespacedv1beta1.SchemeGroupVersion.String()
	mSpec.ObjectMeta = metav1.ObjectMeta{
		Name:        pc.GetName(),
		Labels:      pc.GetLabels(),
		Annotations: pc.GetAnnotations(),
	}
	return &mSpec, err
}

func enrichLocalSecretRefs(pc *namespacedv1beta1.ProviderConfig, mg xpresource.Managed) {
	if pc == nil {
		return
	}
	if pc.Spec.Credentials.SecretRef != nil {
		pc.Spec.Credentials.SecretRef.Namespace = mg.GetNamespace()
	}
	if pc.Spec.Credentials.Upbound != nil &&
		pc.Spec.Credentials.Upbound.WebIdentity != nil &&
		pc.Spec.Credentials.Upbound.WebIdentity.TokenConfig != nil &&
		pc.Spec.Credentials.Upbound.WebIdentity.TokenConfig.SecretRef != nil {
		pc.Spec.Credentials.Upbound.WebIdentity.TokenConfig.SecretRef.Namespace = mg.GetNamespace()
	}
	if pc.Spec.Credentials.WebIdentity != nil &&
		pc.Spec.Credentials.WebIdentity.TokenConfig != nil &&
		pc.Spec.Credentials.WebIdentity.TokenConfig.SecretRef != nil {
		pc.Spec.Credentials.WebIdentity.TokenConfig.SecretRef.Namespace = mg.GetNamespace()
	}
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg xpresource.Managed) (*namespacedv1beta1.ClusterProviderConfig, error) {
	switch managed := mg.(type) {
	case xpresource.LegacyManaged:
		return resolveProviderConfigLegacy(ctx, crClient, managed)
	case xpresource.ModernManaged:
		return resolveProviderConfigModern(ctx, crClient, managed)
	default:
		return nil, errors.New("resource is not a managed")
	}
}

func resolveProviderConfigLegacy(ctx context.Context, client client.Client, mg xpresource.LegacyManaged) (*namespacedv1beta1.ClusterProviderConfig, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := xpresource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return legacyToModernProviderConfigSpec(pc)
}

func resolveProviderConfigModern(ctx context.Context, crClient client.Client, mg xpresource.ModernManaged) (*namespacedv1beta1.ClusterProviderConfig, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrapf(err, "referenced provider config kind %q is invalid for %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}
	pcObj, ok := pcRuntimeObj.(xpresource.ProviderConfig)
	if !ok {
		return nil, errors.Errorf("referenced provider config kind %q is not a provider config type %s/%s", configRef.Kind, mg.GetNamespace(), mg.GetName())
	}

	// Namespace will be ignored if the PC is a cluster-scoped type
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var effectivePC *namespacedv1beta1.ClusterProviderConfig
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		enrichLocalSecretRefs(pc, mg)
		effectivePC = &namespacedv1beta1.ClusterProviderConfig{
			TypeMeta: metav1.TypeMeta{
				APIVersion: namespacedv1beta1.SchemeGroupVersion.String(),
				Kind:       namespacedv1beta1.ClusterProviderConfigKind,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:        pc.GetName(),
				Labels:      pc.GetLabels(),
				Annotations: pc.GetAnnotations(),
			},
			Spec: pc.Spec,
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		// noop
		effectivePC = pc
	default:
		return nil, errors.New("unknown")
	}
	t := xpresource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return effectivePC, nil
}
