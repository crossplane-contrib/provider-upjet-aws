// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/diffserver"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-provider-aws/xpprovider"
	"sigs.k8s.io/controller-runtime/pkg/client"

	namespacedv1beta1 "github.com/upbound/provider-aws/v2/apis/namespaced/v1beta1"
)

// DefaultOfflineAccountID stands in for the AWS account ID when the caller
// does not supply one. The diff server cannot ask STS for the real account,
// and several resources template the account ID into their Terraform ID, so a
// placeholder keeps those IDs syntactically valid. It is the same value the
// online setup falls back to when credentials validation is skipped.
const DefaultOfflineAccountID = localstackAccountID

// Placeholder credentials the offline setup configures the AWS client with.
//
// They are what makes the setup credential-free: aws-sdk-go-base only falls
// back to its credential chain - environment variables, the shared credentials
// and config files, IMDS, the container credentials endpoint - when no static
// key is configured. Handing it a static key short-circuits that chain, so the
// diff server neither reads the host's credentials nor discovers them from the
// environment it runs in. Were the chain to run, a host that has credentials
// would compute diffs against the real AWS account and one that has none would
// fail to configure the client at all.
//
// No request is ever signed with them, because offlineAPIOptions blocks every
// outbound request before it is sent.
const (
	offlineAccessKeyID     = "AKIAOFFLINEDIFFSRVR0" //nolint:gosec // a placeholder, never used to sign a request
	offlineSecretAccessKey = "offlineDiffServerPlaceholderSecretKey000"
)

const (
	errOfflineResolveProviderConfig = "cannot resolve the provider config"
	errOfflineRegion                = "cannot get the region of the managed resource"
	errOfflinePartition             = "cannot configure the AWS partition"
	errOfflineConfigureClient       = "cannot configure the AWS client"
	errOfflineGuardOutbound         = "cannot install the offline egress guard: the Terraform setup's Meta is not an AWS client"

	// fmtErrOfflineRequestBlocked reports an AWS API call that the offline diff
	// server refused to send, named by the service and operation it would have
	// called.
	fmtErrOfflineRequestBlocked = "the offline diff server blocked an outbound AWS request (%s %s): computing this diff requires calling AWS, which is not possible offline"
)

// OfflineTerraformSetupBuilder returns a Terraform setup that configures the
// AWS Terraform provider without credentials and without contacting AWS.
//
// Use it in place of SelectTerraformSetup in the diff server only: the
// resulting terraform.Setup carries an AWS client built from placeholder
// credentials, which is enough for the provider's schema and CustomizeDiff
// logic, and an egress guard that refuses every request the client would send.
//
// The ProviderConfig is still resolved, so that the configuration which
// affects a diff is honoured - the region, the partition, custom endpoints,
// the S3 addressing style - and only the credentials are ignored. Unlike the
// online setup, nothing here resolves credentials, which keeps the setup free
// of the IMDS lookups, token file reads and STS calls that an IRSA, web
// identity or role chain configuration would otherwise perform. The
// placeholders are what guarantee that: see offlineAccessKeyID.
//
// The setup is therefore independent of the host it runs on. It neither
// requires AWS credentials to be present nor uses them when they are, so the
// same request yields the same plan on a developer's machine, in CI, and in a
// cluster.
//
// accountID stands in for the account the credentials would have belonged to.
// It is empty when the caller does not know it, in which case
// DefaultOfflineAccountID is used.
func OfflineTerraformSetupBuilder(config *SetupConfig, accountID string) terraform.SetupFn {
	if accountID == "" {
		accountID = DefaultOfflineAccountID
	}
	return func(ctx context.Context, c client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		pc, err := resolveProviderConfig(ctx, c, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, errOfflineResolveProviderConfig)
		}

		region, err := offlineRegion(mg, pc)
		if err != nil {
			return terraform.Setup{}, err
		}
		// A local copy, so that forcing the metadata API off below is confined
		// to this request rather than applied to the object the in-memory
		// client holds. Nothing offline may consult IMDS: the static
		// credentials already keep it out of the credential chain, and this
		// keeps it out of the rest of the client's configuration too.
		offlinePC := *pc
		offlinePC.Spec.SkipMetadataApiCheck = true

		ps := terraform.Setup{
			// Several external name configurations template the region into
			// the Terraform ID, so it has to be in the configuration and not
			// only in the AWS client.
			Configuration: map[string]any{
				keyRegion: region,
			},
			ClientMetadata: map[string]string{
				keyAccountID: accountID,
				keyPartition: partitionAWS,
			},
		}
		// Resolving the partition is pure computation over the region and the
		// ProviderConfig, so the offline setup can do it exactly as the online
		// one does.
		if err := setPartition(&aws.Config{Region: region}, &offlinePC, &ps); err != nil {
			return terraform.Setup{}, errors.Wrap(err, errOfflinePartition)
		}

		if config.TerraformProvider == nil {
			return terraform.Setup{}, errors.New("terraform provider cannot be nil")
		}
		// The AWS client is fully built, so that the provider can consult its
		// schema and run its CustomizeDiff functions, but it is built from
		// placeholder credentials and cannot authenticate against anything.
		creds := aws.Credentials{
			AccessKeyID:     offlineAccessKeyID,
			SecretAccessKey: offlineSecretAccessKey,
		}
		if err := configureNoForkAWSClient(ctx, &ps, config, region, creds, &offlinePC); err != nil {
			return terraform.Setup{}, errors.Wrap(err, errOfflineConfigureClient)
		}
		if err := guardOutboundRequests(ps.Meta); err != nil {
			return terraform.Setup{}, err
		}
		return ps, nil
	}
}

// guardOutboundRequests installs the egress guard on the AWS client the
// Terraform setup carries. Offline diffs depend on it, so a client it cannot be
// installed on is an error rather than a degradation.
func guardOutboundRequests(meta any) error {
	c, ok := meta.(*xpprovider.AWSClient)
	if !ok {
		return errors.New(errOfflineGuardOutbound)
	}
	// Every service client the Terraform provider asks for is built from the
	// client's AWS config, and none has been built yet, so appending the option
	// here covers all of them.
	c.AppendAPIOptions(denyOutboundRequests)
	return nil
}

// denyOutboundRequests refuses every AWS API call, failing the diff
// immediately instead of sending a request the diff server has no business
// sending. A handful of resources call AWS from their CustomizeDiff function -
// aws_sfn_state_machine validates its definition, aws_cloudcontrolapi_resource
// reads its CloudFormation type, aws_snapshot_create_volume_permission reads
// the snapshot - and their diffs simply cannot be computed offline. Without
// this guard such a call would leave the process, be retried, and finally fail
// with a connection or authentication error that says nothing about why.
//
// The error is a diff-computation-not-supported error, which the diff server
// reports as a FailedPrecondition carrying a DIFF_COMPUTATION_NOT_SUPPORTED
// violation, so that a client can fall back to diffing the manifests.
func denyOutboundRequests(stack *middleware.Stack) error {
	// The guard sits at the very front of the Initialize step, ahead of
	// endpoint resolution, signing and the retry loop, so that it runs exactly
	// once per operation and nothing is attempted before it. The service and
	// operation names are on the context before the stack runs, so they are
	// available this early.
	return stack.Initialize.Add(middleware.InitializeMiddlewareFunc("offlineDiffEgressGuard",
		func(ctx context.Context, _ middleware.InitializeInput, _ middleware.InitializeHandler) (middleware.InitializeOutput, middleware.Metadata, error) {
			return middleware.InitializeOutput{}, middleware.Metadata{}, diffserver.NewDiffComputationNotSupportedError(
				errors.Errorf(fmtErrOfflineRequestBlocked, awsmiddleware.GetServiceID(ctx), awsmiddleware.GetOperationName(ctx)))
		}), middleware.Before)
}

// offlineRegion returns the region to configure the Terraform provider with,
// without resolving credentials. It is the region the managed resource
// declares, falling back to a partition-appropriate default for the resources
// that are global and therefore declare none. This mirrors what
// getAWSConfigWithDefaultRegion arrives at online.
func offlineRegion(mg xpresource.Managed, pc *namespacedv1beta1.ClusterProviderConfig) (string, error) {
	region, err := getRegion(mg)
	if err != nil {
		return "", errors.Wrap(err, errOfflineRegion)
	}
	if region != "" {
		return region, nil
	}
	// A global resource has no region of its own, but the Terraform AWS
	// provider still requires a non-empty one.
	gvk := mg.GetObjectKind().GroupVersionKind()
	return getGlobalRegion(gvk.Group, gvk.Kind, pc), nil
}
