/*
Copyright 2021 Upbound Inc.
*/

package rds

import (
	"context"
	"fmt"
	"github.com/upbound/upjet/pkg/types/comments"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/password"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for rds group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_rds_cluster", func(r *config.Resource) {
		// Mutually exclusive with aws_rds_cluster_role_association
		config.MoveToStatus(r.TerraformResource, "iam_roles")
		r.References["s3_import.bucket_name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["restore_to_point_in_time.source_cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.References["db_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["endpoint"].(string); ok {
				conn["endpoint"] = []byte(a)
			}
			if a, ok := attr["reader_endpoint"].(string); ok {
				conn["reader_endpoint"] = []byte(a)
			}
			return conn, nil
		}
	})

	p.AddResourceConfigurator("aws_rds_cluster_instance", func(r *config.Resource) {
		r.References["restore_to_point_in_time.source_db_instance_identifier"] = config.Reference{
			Type: "Instance",
		}
		r.References["s3_import.bucket_name"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket",
		}
		r.References["kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["performance_insights_kms_key_id"] = config.Reference{
			Type: "github.com/upbound/provider-aws/apis/kms/v1beta1.Key",
		}
		r.References["restore_to_point_in_time.source_cluster_identifier"] = config.Reference{
			Type: "Cluster",
		}
		r.References["security_group_names"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup",
			RefFieldName:      "SecurityGroupNameRefs",
			SelectorFieldName: "SecurityGroupNameSelector",
		}
		r.References["parameter_group_name"] = config.Reference{
			Type: "ParameterGroup",
		}
		r.References["db_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		delete(r.References, "engine")
		delete(r.References, "engine_version")
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_db_instance", func(r *config.Resource) {
		r.References["db_subnet_group_name"] = config.Reference{
			Type: "SubnetGroup",
		}
		r.References["kms_key_id"] = config.Reference{
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}
		r.UseAsync = true
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"name", "db_name"},
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["endpoint"].(string); ok {
				conn["endpoint"] = []byte(a)
			}
			if a, ok := attr["address"].(string); ok {
				conn["address"] = []byte(a)
				conn["host"] = []byte(a)
			}
			if a, ok := attr["username"].(string); ok {
				conn["username"] = []byte(a)
			}
			if a, ok := attr["port"]; ok {
				conn["port"] = []byte(fmt.Sprintf("%v", a))
			}

			return conn, nil
		}
		desc, _ := comments.New("If true, the password will be auto-generated and"+
			" stored in the Secret referenced by the passwordSecretRef field.",
			comments.WithTFTag("-"))
		r.TerraformResource.Schema["auto_generate_password"] = &schema.Schema{
			Type:        schema.TypeBool,
			Optional:    true,
			Description: desc.String(),
		}
		r.InitializerFns = append(r.InitializerFns, func(client client.Client) managed.Initializer {
			return managed.InitializerFn(func(ctx context.Context, mg resource.Managed) error {
				paved, err := fieldpath.PaveObject(mg)
				if err != nil {
					return err
				}
				sel := &v1.SecretKeySelector{}
				if err := paved.GetValueInto("spec.forProvider.passwordSecretRef", sel); err != nil {
					return errors.Wrap(err, "cannot unmarshal passwordSecretRef into a secret key selector")
				}
				s := &corev1.Secret{}
				if err := client.Get(ctx, types.NamespacedName{Namespace: sel.Namespace, Name: sel.Name}, s); resource.IgnoreNotFound(err) != nil {
					return errors.Wrap(err, "cannot get password secret")
				}
				if err == nil && len(s.Data[sel.Key]) != 0 {
					// Password is already set.
					return nil
				}
				// At this point, either the secret doesn't exist, or it doesn't
				// have the password filled.
				gen, err := paved.GetBool("spec.forProvider.autoGeneratePassword")
				if err != nil {
					return errors.Wrap(err, "cannot get autoGeneratePassword field value")
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
		})
		r.TerraformResource.Schema["password"].Description = "Password for the master DB user. If you set autoGeneratePassword to true, the Secret referenced here will be created or updated with generated password if it does not already contain one."
	})

	p.AddResourceConfigurator("aws_db_proxy", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_proxy_endpoint", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_rds_cluster_activity_stream", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_snapshot", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_db_option_group", func(r *config.Resource) {
		delete(r.References, "option.option_settings.value")
	})

	p.AddResourceConfigurator("aws_db_proxy_target", func(r *config.Resource) {
		delete(r.References, "target_group_name")
	})

	p.AddResourceConfigurator("aws_rds_cluster_endpoint", func(r *config.Resource) {
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_rds_cluster_role_association", func(r *config.Resource) {
		r.UseAsync = true
	})
}
