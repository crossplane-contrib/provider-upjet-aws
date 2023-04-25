/*
Copyright 2021 Upbound Inc.
*/

package elasticache

import (
	"context"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/password"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/comments"

)

const (
	errGetPasswordSecret = "cannot get password secret"
)

// Configure adds configurations for elasticache group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_elasticache_cluster", func(r *config.Resource) {
		r.References = config.References{
			"parameter_group_name": config.Reference{
				Type: "ParameterGroup",
			},
			"subnet_group_name": config.Reference{
				Type: "SubnetGroup",
			},
			"security_group_ids": config.Reference{
				Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
				RefFieldName:      "SecurityGroupIDRefs",
				SelectorFieldName: "SecurityGroupIDSelector",
			},
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["cluster_address"].(string); ok {
				conn["cluster_address"] = []byte(a)
			}
			if a, ok := attr["configuration_endpoint"].(string); ok {
				conn["configuration_endpoint"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_elasticache_replication_group", func(r *config.Resource) {
		r.References["subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.References["security_group_ids"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}
		r.References["kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.LateInitializer = config.LateInitializer{
			// Conflicting configuration arguments: "number_cache_clusters": conflicts with cluster_mode.0.num_node_groups
			IgnoredFields: []string{
				"cluster_mode",
				"num_node_groups",
				"num_cache_clusters",
				"number_cache_clusters",
				"replication_group_description",
				"description",
			},
		}
		delete(r.References, "log_delivery_configuration.destination")
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["configuration_endpoint_address"].(string); ok {
				conn["configuration_endpoint_address"] = []byte(a)
			}
			if a, ok := attr["primary_endpoint_address"].(string); ok {
				conn["primary_endpoint_address"] = []byte(a)
			}
			if a, ok := attr["reader_endpoint_address"].(string); ok {
				conn["reader_endpoint_address"] = []byte(a)
			}
			if a, ok := attr["auth_token"].(string); ok {
				conn["auth_token"] = []byte(a)
			}
			return conn, nil
		}
		desc, _ := comments.New("If true, the auth_token will be auto-generated and"+
			" stored in the Secret referenced by the authTokenSecretRef field.",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["auto_generate_authtoken"] = &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Description: desc.String(),
		}
		r.InitializerFns = append(r.InitializerFns,
			PasswordGenerator(
				"spec.forProvider.authTokenSecretRef",
				"spec.forProvider.autoGenerateAuthToken",
			))
		r.TerraformResource.Schema["auth_token"].Description = "AuthToken for the " +
			"connnection. If you set autoGenerateAuthToken to true, the AuthToken" +
			" referenced here will be created or updated with generated auth_token" +
			" if it does not already contain one."
	})

	p.AddResourceConfigurator("aws_elasticache_user_group", func(r *config.Resource) {
		r.References["user_ids"] = config.Reference{
			Type:              "User",
			RefFieldName:      "UserIDRefs",
			SelectorFieldName: "UserIDSelector",
		}
	})
}

// PasswordGenerator returns an InitializerFn that will generate a password
// for a resource if the toggle field is set to true and the secret referenced
// by the secretRefFieldPath is not found or does not have content corresponding
// to the password key.
func PasswordGenerator(secretRefFieldPath, toggleFieldPath string) config.NewInitializerFn {
	return func(client client.Client) managed.Initializer {
		return managed.InitializerFn(func(ctx context.Context, mg resource.Managed) error {
			paved, err := fieldpath.PaveObject(mg)
			if err != nil {
				return errors.Wrap(err, "cannot pave object")
			}
			sel := &v1.SecretKeySelector{}
			err = paved.GetValueInto(secretRefFieldPath, sel)
			if err != nil {
				return errors.Wrapf(resource.Ignore(fieldpath.IsNotFound, err), "cannot unmarshal %s into a secret key selector", secretRefFieldPath)
			}
			s := &corev1.Secret{}
			if err := client.Get(ctx, types.NamespacedName{Namespace: sel.Namespace, Name: sel.Name}, s); resource.IgnoreNotFound(err) != nil {
				return errors.Wrap(err, errGetPasswordSecret)
			}
			if err == nil && len(s.Data[sel.Key]) != 0 {
				// Password is already set.
				return nil
			}
			// At this point, either the secret doesn't exist, or it doesn't
			// have the password filled.
			gen, err := paved.GetBool(toggleFieldPath)
			if resource.Ignore(fieldpath.IsNotFound, err) != nil {
				return errors.Wrapf(err, "cannot get the value of %s", toggleFieldPath)
			}
			if !gen {
				// Password is not set, and we don't want to generate one.
				return nil
			}
			pw, err := password.Generate()
			if err != nil {
				return errors.Wrap(err, "cannot generate password")
			}
			s.SetName(sel.Name)
			s.SetNamespace(sel.Namespace)
			if s.Data == nil {
				s.Data = make(map[string][]byte, 1)
			}
			s.Data[sel.Key] = []byte(pw)
			return errors.Wrap(resource.NewAPIPatchingApplicator(client).Apply(ctx, s), "cannot apply password secret")
		})
	}
}