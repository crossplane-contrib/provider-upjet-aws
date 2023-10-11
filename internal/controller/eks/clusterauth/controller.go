/*
Copyright 2022 Upbound Inc.
*/

package clusterauth

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	tjcontroller "github.com/crossplane/upjet/pkg/controller"
	ujresource "github.com/crossplane/upjet/pkg/resource"

	"github.com/upbound/provider-aws/apis/eks/v1beta1"
	"github.com/upbound/provider-aws/apis/v1alpha1"
	"github.com/upbound/provider-aws/internal/clients"
	"github.com/upbound/provider-aws/internal/features"
)

const (
	additionalDurationForExpiration = 5 * time.Minute

	errNotClusterAuth  = "managed resource is not a ClusterAuth custom resource"
	errDescribeCluster = "cannot describe cluster"
	errGetKubeconfig   = "cannot get kubeconfig"
	errStatusUpdate    = "cannot update status of ClusterAuth"
)

// Setup adds a controller that reconciles ClusterAuth.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1beta1.ClusterAuth_GroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), v1alpha1.StoreConfigGroupVersionKind))
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.ClusterAuth{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1beta1.ClusterAuth_GroupVersionKind),
			managed.WithExternalConnecter(&connector{
				kube:               mgr.GetClient(),
				newEKSClientFn:     eks.NewFromConfig,
				newPresignClientFn: newPresignClient,
			}),
			// We use a constant poll interval here to make sure we get a chance
			// to refresh the token before it expires.
			managed.WithPollInterval(time.Minute*1),
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
			managed.WithConnectionPublishers(cps...)))
}

type connector struct {
	kube               client.Client
	newEKSClientFn     func(cfg aws.Config, optFns ...func(*eks.Options)) *eks.Client
	newPresignClientFn func(cfg aws.Config, optFns ...func(*sts.Options)) *sts.PresignClient
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cfg, err := clients.GetAWSConfig(ctx, c.kube, mg)
	if err != nil {
		return nil, err
	}
	return &external{
			eksClient:     c.newEKSClientFn(*cfg),
			presignClient: c.newPresignClientFn(*cfg),
			kube:          c.kube},
		nil
}

type external struct {
	eksClient     *eks.Client
	presignClient *sts.PresignClient
	kube          client.Client
}

func (e *external) Observe(_ context.Context, mg resource.Managed) (managed.ExternalObservation, error) { // nolint:gocyclo
	cr, ok := mg.(*v1beta1.ClusterAuth)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotClusterAuth)
	}
	if meta.WasDeleted(cr) {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	if cr.Status.AtProvider.LastRefreshTime == nil {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	deadline := cr.Status.AtProvider.LastRefreshTime.Add(cr.Spec.ForProvider.RefreshPeriod.Duration)
	if time.Now().After(deadline) {
		return managed.ExternalObservation{
			ResourceExists:   true,
			ResourceUpToDate: false,
		}, nil
	}
	cr.Status.SetConditions(xpv1.Available())
	ujresource.SetUpToDateCondition(mg, true)
	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1beta1.ClusterAuth)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotClusterAuth)
	}
	cl, err := e.eksClient.DescribeCluster(ctx, &eks.DescribeClusterInput{Name: aws.String(cr.Spec.ForProvider.ClusterName)})
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errDescribeCluster)
	}
	if aws.ToString(cl.Cluster.CertificateAuthority.Data) == "" {
		return managed.ExternalCreation{}, errors.New("ca data from the retrieved cluster is empty")
	}
	// NOTE(muvaf): The maximum time allowed for a token to live is 15 minutes
	// even though API allows setting longer durations. Additional duration is
	// add cushion so that we have the room for reconciliation to kick in at most
	// in 5 minutes.
	d := cr.Spec.ForProvider.RefreshPeriod.Duration + additionalDurationForExpiration
	if d > time.Minute*15 {
		d = time.Minute * 15
	}
	conn, err := GetConnectionDetails(
		ctx,
		e.presignClient,
		cl.Cluster,
		d,
	)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errGetKubeconfig)
	}
	t := metav1.NewTime(time.Now())
	cr.Status.AtProvider.LastRefreshTime = &t
	cr.Status.SetConditions(xpv1.Available())
	// NOTE(muvaf): We need to update status by ourselves because after-math
	// of Create doesn't include updating the status, hence the lastRefreshTime
	// gets lost.
	if err := e.kube.Status().Update(ctx, cr); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errStatusUpdate)
	}
	return managed.ExternalCreation{ConnectionDetails: conn}, nil
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1beta1.ClusterAuth)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotClusterAuth)
	}
	cl, err := e.eksClient.DescribeCluster(ctx, &eks.DescribeClusterInput{Name: aws.String(cr.Spec.ForProvider.ClusterName)})
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errDescribeCluster)
	}
	if aws.ToString(cl.Cluster.CertificateAuthority.Data) == "" {
		return managed.ExternalUpdate{}, errors.New("ca data from the retrieved cluster is empty")
	}
	// NOTE(muvaf): The maximum time allowed for a token to live is 15 minutes
	// even though API allows setting longer durations. Additional duration is
	// add cushion so that we have the room for reconciliation to kick in at most
	// in 5 minutes.
	d := cr.Spec.ForProvider.RefreshPeriod.Duration + additionalDurationForExpiration
	if d > time.Minute*15 {
		d = time.Minute * 15
	}
	conn, err := GetConnectionDetails(
		ctx,
		e.presignClient,
		cl.Cluster,
		d,
	)
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errGetKubeconfig)
	}
	t := metav1.NewTime(time.Now())
	cr.Status.AtProvider.LastRefreshTime = &t
	return managed.ExternalUpdate{ConnectionDetails: conn}, nil
}

func (e *external) Delete(_ context.Context, _ resource.Managed) error {
	return nil
}
