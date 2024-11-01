# AWS ProviderConfig E2E testing

## Introduction

This Crossplane configuration package aims to provide a base environment for testing 
various `ProviderConfig` scenarios of `provider-upjet-aws`.

It provisions:
- An AWS EKS cluster
- AWS IRSA-related IAM resources for the EKS cluster for testing IRSA authentication
- AWS WebIdentity authentication related IAM resources for the EKS cluster, to test WebIdentity authentication
After creating the EKS cluster, it deploys:
- `Crossplane` into the EKS cluster
- `DeploymentRuntimeConfig`s for XP providers to enable IRSA
- `provider-family-aws` and 2 example providers for testing
  - `provider-aws-ec2` and `provider-aws-rds`
- various AWS `ProviderConfig` manifests for testing scenarios
- example Managed Resources (MRs) from AWS EC2 and RDS groups, referencing the `ProviderConfig`s in test 

## Package Structure

The package consists of the Composite Resource (XR) 
`xe2etestclusters.aws.platformref.upbound.io`

This composite resource makes use of the existing configuration 
packages `configuration-aws-eks` and `configuration-aws-eks-irsa`.

It is structured in a way that starting from a local crossplane control plane, 
it sets up another control plane in an EKS cluster with Crossplane. 
Via the `provider-kubernetes` and `provider-helm` at the local control plane, 
the remote EKS control plane is bootstrapped and relevant test resources are deployed.
This can be considered as a "A crossplane control plane is Managed by another Crossplane control plane".
This setup allows conducting tests from a local control plane.

![pc-e2e-diagram.png](docs%2Fimg%2Fpc-e2e-diagram.png)

Explicit deletion ordering inside composition is implemented via `Usages`. For some resources, 
Crossplane runtime already handles the implicit dependencies such as MR <-> ProviderConfig, ProviderConfig <-> Providers.
The dependencies are depicted in the diagram above.

### `e2etestcluster.platformref.upbound.io` XRC

You can find an example test cluster claim at [package/examples/e2etestcluster-claim.yaml](package%2Fexamples%2Fe2etestcluster-claim.yaml)
When this claim is is created and ready, it means that the tests are passing.

```yaml
apiVersion: aws.platformref.upbound.io/v1alpha1
kind: E2ETestCluster
metadata:
  name: aws-pc-e2e-test
  namespace: default
spec:
  compositeDeletePolicy: Foreground
  parameters:
    id: aws-pc-e2e-test
    region: us-west-2 # EKS cluster region
    version: "1.29" # EKS cluster k8s version
    iam:
      # replace with your custom roleArn that will administer the EKS cluster:
      roleArn: "arn:aws:iam::123456789012:role/mydefaulteksadminrole"
    nodes: # eks nodes configuration
      count: 1  
      instanceType: t3.medium
    irsa: # IRSA configuration for the AWS role that will be used by XP providers
      condition: StringEquals
      # The policy of the IRSA role
      policyDocument: |
        {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Action": [
                        "ec2:*"
                    ],
                    "Effect": "Allow",
                    "Resource": "*"
                }
            ]
        } 
      serviceAccount: # name of the k8s service account to be created for provider pods 
        name: my-xpsa
        namespace: upbound-system
    targetClusterParameters: # the parameters for the target EKS control plane cluster 
      provider: # provider package urls to be used in testing
        familyPackage: "xpkg.upbound.io/upbound/provider-family-aws:v1.16.0"
        ec2Package: "xpkg.upbound.io/upbound/provider-aws-ec2:v1.16.0"
        rdsPackage: "xpkg.upbound.io/upbound/provider-aws-rds:v1.16.0"
        kafkaPackage: "xpkg.upbound.io/upbound/provider-aws-kafka:v1.16.0"
      crossplaneVersion: 1.17.2 # the crossplane version to be installed in the testing control plane
  writeConnectionSecretToRef:
    name: aws-pc-e2e-test-kubeconfig
status:
  irsa:
    roleArn: irsa-role-arn
    chainedRoleARNs:
      - "chained-role-arn"
  webIdentity:
    roleArn: webid-role-arn
    chainedRoleARNs:
      - "chained-role-arn"

```

## Usage

### Prerequisites

- An AWS account and relevant credentials capable of creating and managing EC2, EKS and IAM resources
- An OCI image registry from which the EKS cluster can pull images (e.g. Dockerhub, xpkg.upbound.io)

### Utilizing Uptest

In order to conduct an e2e test using uptest, the `make` targets can be used.

### option 1. inside configuration package: make target `e2e`

This make target:
- builds the `providerconfig-aws-e2e-test` configuration package
- spins up a local `kind` cluster
- deploys the `providerconfig-aws-e2e-test` configuration package to the `kind` cluster
- runs the e2e tests

This make target expects the target AWS provider packages in test to be already built and pushed to
a registry that the target EKS cluster can reach.  

The make target expects the following environment variables to be set:

- `AWS_FAMILY_PACKAGE_IMAGE`: The package URL for `provider-family-aws` 
- `AWS_EC2_PACKAGE_IMAGE`: The package URL for `provider-aws-ec2`
- `AWS_RDS_PACKAGE_IMAGE`: The package URL for `provider-aws-rds`
- `AWS_KAFKA_PACKAGE_IMAGE`: The package URL for `provider-aws-kafka`
- `AWS_EKS_IAM_DEFAULT_ADMIN_ROLE`: the ARN of an existing IAM role. This will be assigned as the E2E test EKS cluster default admin
- `TARGET_CROSSPLANE_VERSION`: The target crossplane version to be deployed into the testing cluster
- `UPTEST_CLOUD_CREDENTIALS`: The AWS credentials for the AWS account that the e2e tests will run on. Should be in the format of AWS CLI INI config. 

An example usage:

my-aws-creds.txt
```ini
[default]
aws_access_key_id = YOUR-AWS-ACCESS-KEY
aws_secret_access_key = your-aws-secret-access-key
```

```shell
export AWS_FAMILY_PACKAGE_IMAGE="xpkg.upbound.io/upbound/provider-family-aws:v1.16.0"
export AWS_EC2_PACKAGE_IMAGE="xpkg.upbound.io/upbound/provider-aws-ec2:v1.16.0"
export AWS_RDS_PACKAGE_IMAGE="xpkg.upbound.io/upbound/provider-aws-rds:v1.16.0"
export AWS_KAFKA_PACKAGE_IMAGE="xpkg.upbound.io/upbound/provider-aws-kafka:v1.16.0"
export AWS_EKS_IAM_DEFAULT_ADMIN_ROLE="arn:aws:iam::123456789012:role/mydefaulteksadminrole"
export TARGET_CROSSPLANE_VERSION="1.17.2"
export UPTEST_CLOUD_CREDENTIALS="$(cat my-aws-creds.txt)"
# from repo root
make -C e2e/providerconfig-aws-e2e-test e2e
```

### option 2. with provider image publish: target `providerconfig-e2e`

This make target:
- builds and publishes the providers
- builds the `providerconfig-aws-e2e-test` configuration package
- spins up a local `kind` cluster
- deploys the `providerconfig-aws-e2e-test` configuration package to the `kind` cluster
- runs the e2e tests using `uptest` with the published provider images

The make target expects 
- `XPKG_REG_ORGS`: the target OCI repository URL for provider images to be published
- `VERSION`: the version tag of the published provider images
- `UPTEST_CLOUD_CREDENTIALS`: The AWS credentials for the AWS account that the e2e tests will run on. Should be in the format of AWS CLI INI config.
- `AWS_EKS_IAM_DEFAULT_ADMIN_ROLE`: the ARN of an existing IAM role. This will be assigned as the E2E test EKS cluster default admin

example usage:
```shell
export UPTEST_CLOUD_CREDENTIALS="$(cat my-aws-creds.txt)"
export AWS_EKS_IAM_DEFAULT_ADMIN_ROLE="arn:aws:iam::123456789012:role/mydefaulteksadminrole"
# from repo root
make VERSION=v1.4.0-testversion XPKG_REG_ORGS=index.docker.io/erhancag providerconfig-e2e
```

### Via Github Actions from PR 
TBD

### Manual
In your desired k8s environment:

- Deploy Crossplane
- build and publish the e2e testing configuration package
- install configuration package
- create the example claim in the `package/examples/e2etestcluster-claim.yaml`. Before creation, modify the claim accordingly if needed.
- wait for the claim to be ready.

## Extending Test Scenarios
TBD