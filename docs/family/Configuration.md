---
title: Configuration
weight: 2
---

# AWS official provider-family documentation
Upbound supports and maintains the Upbound AWS official provider-family.

## Install the provider-family-aws
### Prerequisites
#### Upbound Up command-line
The Upbound Up command-line simplifies configuration and management of Upbound
Universal Crossplane (UXP) and interacts with the Upbound Marketplace to manage
users and accounts.

Install `up` with the command:
```shell
curl -sL "https://cli.upbound.io" | sh
```
More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

#### Upbound Universal Crossplane
UXP is the Upbound official enterprise-grade distribution of Crossplane for
self-hosted control planes. 

Install UXP into your Kubernetes cluster using the Up command-line.

```shell
up uxp install
```

Find more information in the [Upbound UXP
documentation](https://docs.upbound.io/uxp/).

### Install the provider-family-aws

Install the Upbound official AWS provider-famil with the following configuration file.
For instance, let's install the `provider-aws-s3`

_Note_: The first provider installed of a family also installs an extra provider-family Provider.
The provider-family provider manages the ProviderConfig for all other providers in the same family.

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws-s3
spec:
  package: xpkg.upbound.io/upbound/provider-aws-s3:<version>
EOF
```

Define the `provider-aws-s3` version with `spec.package`.

Install the `provider-aws-s3` with `kubectl apply -f`.

Verify the configuration with `kubectl get providers`.

```shell
NAME                          INSTALLED   HEALTHY   PACKAGE                                               AGE
provider-aws-s3               True        True      xpkg.upbound.io/upbound/provider-aws-s3:v0.40.0       6m39s
upbound-provider-family-aws   True        True      xpkg.upbound.io/upbound/provider-family-aws:v0.40.0   6m30s
```

View the Crossplane [Provider CRD
definition](https://doc.crds.dev/github.com/crossplane/crossplane/pkg.crossplane.io/Provider/v1)
to view all available `Provider` options.

## Configure the provider-family-aws
The AWS provider-family requires credentials for authentication to AWS. The AWS
provider-family consumes the credentials from a Kubernetes secret object.

### Configure authentication
Upbound supports authentication to AWS via access keys, service accounts or with
`AssumeRole`.

Apply the specific authentication method with a `ProviderConfig` object, applied
to the `Provider`.

#### Authenticate using AWS access keys
Authenticating with AWS access keys requires creating a Kubernetes secret object
and storing the AWS keys inside Kubernetes.

##### Place the AWS access keys in a text file
Create a text file containing the AWS account `aws_access_key_id` and
`aws_secret_access_key`. The [AWS
documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-creds)
provides information on how to generate these keys.

```ini
[default]
aws_access_key_id = <aws_access_key>
aws_secret_access_key = <aws_secret_key>
```

More information about AWS credential files are in the [AWS credential file
documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html).

##### Generate a Kubernetes secret with the AWS access keys
Create the secret with the command  

```shell
kubectl create secret generic \
<secret name> \
--from-file=key-file=<aws_credentials_file.txt>
```

For example, to create a secret named `aws-secret` from a text file named
`aws-credentials.txt`
```shell
$ kubectl create secret generic aws-secret --from-file=key-file=aws-credentials.txt
$ kubectl describe secret aws-secret
Name:         aws-secret
Namespace:    default
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
key-file:  116 bytes
```

##### Create a ProviderConfig object
Apply the secret in a `ProviderConfig` Kubernetes configuration file. For
example using a secret named `aws-secret`:

```yaml
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
    name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: default
      name: aws-secret
      key: key-file
```

The `spec.secretRef` describes the parameters of the secret to use. 
* `namespace` is the Kubernetes namespace the secret is in.
* `name` is the name of the Kubernetes `secret` object.
* `key` is the `Data` field from `kubectl describe secret`.

View the [ProviderConfig
CRD](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/resources/aws.upbound.io/ProviderConfig/v1beta1)
definition to view all available `ProviderConfig` options.

#### Authenticate using IAM Roles for Service Accounts
Universal Crossplane clusters running inside Amazon Elastic Kubernetes Service
(`EKS`) can use [IAM Roles for Service Accounts
(`IRSA`)](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html)
to authenticate the AWS provider.

An IRSA configuration requires multiple components:
* Enable the [IAM OIDC
  provider](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html)
  for EKS.
* Creating an IAM policy granting the AWS provider access to AWS resources.
* Creating an IAM role for the AWS provider to associate with the AWS provider.
* Creating a Kubernetes service account.
* Create a `ControllerConfig` to associate the IAM role ARN.
* Apply the `ControllerConfig` to the `Provider`.
* Instruct the `ProviderConfig` to use `IRSA` credentials.

<!-- Disable heading acronym rule to ignore "OIDC" -->
<!-- vale Microsoft.HeadingAcronyms = NO -->
##### Enable an IAM OIDC provider
<!-- vale Microsoft.HeadingAcronyms = YES -->
The EKS cluster must have an IAM OpenID Connect (`OIDC`) provider enabled to
configure IRSA. The AWS documentation contains full details on [enabling IAM
OIDC
providers](https://docs.aws.amazon.com/eks/latest/userguide/enable-iam-roles-for-service-accounts.html).

Using the [`eksctl`](https://eksctl.io/) tool, create an IAM OIDC provider with
the command

```shell
eksctl utils associate-iam-oidc-provider \ 
--cluster <EKS cluster name> \
--approve
```

Confirm IAM OIDC provider creation using the [AWS
command-line](https://aws.amazon.com/cli/)
```shell
$ aws iam list-open-id-connect-providers
{
    "OpenIDConnectProviderList": [
        {
            "Arn": "arn:aws:iam::000000000000:oidc-provider/oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4"
        }
    ]
}
```

##### Create an IAM policy
Define the actions the AWS provider can take by creating an IAM policy. 

For example, here is a custom IAM policy to enable `SystemAdministrator` level
access. 
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "*",
            "Resource": "*"
        }
    ]
}
```
Apply the policy using the AWS command-line command 
```shell
aws iam create-policy \
--policy-name <new policy name> \
--policy-document file://<policy json file>
```

For example, to create a new policy named `custom-irsa-policy` from a policy
file named `custom-policy.json`:  
```shell
$ aws iam create-policy --policy-name custom-irsa-policy --policy-document file://custom-policy.json
{
    "Policy": {
        "PolicyName": "custom-irsa-policy",
        "PolicyId": "ANPAZBZV2IPHDUBU5BF56",
        "Arn": "arn:aws:iam::000000000000:policy/custom-irsa-policy",
        "Path": "/",
        "DefaultVersionId": "v1",
        "AttachmentCount": 0,
        "PermissionsBoundaryUsageCount": 0,
        "IsAttachable": true,
        "CreateDate": "2022-08-17T20:35:22+00:00",
        "UpdateDate": "2022-08-17T20:35:22+00:00"
    }
}
```

_Note:_ if you plan to use an AWS managed policy, a custom policy isn't
required.

##### Create an IAM role
Creating a Kubernetes service account requires an IAM role to apply an IAM
policy. The AWS documentation contains full details on [creating a role and
assigning it to the Kubernetes service
account](https://docs.aws.amazon.com/eks/latest/userguide/associate-service-account-role.html).

_Note:_ `eksctl` is the simplest way to create the required role, attach the
policy and generate a service account. You **must** change the role trust policy
after creating it with `eksctl`.

To use `eksctl` to create a new service account and IAM role use the command
`eksctl create iamserviceaccount`. 

```shell
eksctl create iamserviceaccount \
--name <kubernetes service account name> \
--role-name <IAM role name> \
--cluster <the name of the EKS cluster> 
--attach-policy-arn <the ARN of the policy to apply> \
--namespace upbound-system \
--approve
```

| Configuration option | Description | 
| ---- | ---- |
| `--name` | The name of the Kubernetes service account to create. |
| `--role-name` | The name of the AWS IAM role to create. | 
| `--cluster` | The name of the EKS cluster. |
| `--attach-policy-arn` | The ARN of the policy to attach to this service account and role. | 
| `--namespace` | The namespace to create the Kubernetes service account in. This must be the same namespace as Universal Crossplane. (The Universal Crossplane default namespace is `upbound-system`.) |

For example, to create a new service account with the configuration:

| Configuration option | Configuration value | 
| ---- | ---- |
| `--name` | `my-upbound-sa` |
| `--role-name` | `eks-test-role` | 
| `--cluster` | `upbound-docs` |
| `--attach-policy-arn` | `arn:aws:iam::000000000000:policy/custom-irsa-policy` |
| `--namespace` | `upbound-system` |

Use the command
```shell
eksctl create iamserviceaccount \
--name my-upbound-sa \
--role-name eks-test-role \
--cluster upbound-docs \
--attach-policy-arn arn:aws:iam::000000000000:policy/custom-irsa-policy \
--namespace upbound-system \
--approve
```

_Note:_ the policy ARN value comes from the `Arn` field from the command `aws
iam create-policy`. To find the ARN of the policy use the command `aws iam
list-policies` replacing `POLICY_NAME` with the name of the policy.

```shell
$ aws iam list-policies --query 'Policies[?PolicyName==`POLICY_NAME`].Arn' --output text`
arn:aws:iam::000000000000:policy/custom-irsa-policy
```

Verify the creation of the service account with the command `kubectl describe sa
-n <namespace> <service account name>`. The `Annotations` field is the newly
created IAM role.

From the example service account named `my-upbound-sa`:
```yaml
$ kubectl describe sa \
-n upbound-system \
my-upbound-sa
Name:                my-upbound-sa
Namespace:           upbound-system
Labels:              app.kubernetes.io/managed-by=eksctl
Annotations:         eks.amazonaws.com/role-arn: arn:aws:iam::000000000000:role/eks-test-role
Image pull secrets:  <none>
Mountable secrets:   my-upbound-sa-token-spq5k
Tokens:              my-upbound-sa-token-spq5k
Events:              <none>
```

Confirm the attachment between the IAM policy and new IAM role with the command 
```shell
aws iam list-attached-role-policies \
--role-name <role name> \
--query AttachedPolicies[].PolicyArn \
--output text
```

For example,
```shell
$ aws iam list-attached-role-policies \
--role-name eks-test-role \ 
--query AttachedPolicies[].PolicyArn \
--output text
arn:aws:iam::000000000000:policy/custom-irsa-policy
```
The output of the command matches the policy ARN.


Next, verify the new IAM role with the command `aws iam get-role --role-name
<role name> --query Role.AssumeRolePolicyDocument`.

Using the example role name `eks-test-role`
```shell
$ aws iam get-role \
--role-name eks-test-role \
--query Role.AssumeRolePolicyDocument
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Federated": "arn:aws:iam::000000000000:oidc-provider/oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4"
            },
            "Action": "sts:AssumeRoleWithWebIdentity",
            "Condition": {
                "StringEquals": {
                    "oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4:aud": "sts.amazonaws.com",
                    "oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4:sub": "system:serviceaccount:upbound-system:my-upbound-sa"
                }
            }
        }
    ]
}
```
##### Update the IAM role
The IAM role created by `eksctl` doesn't have the correct `Conditions` for the
AWS provider. 

Update the role `Trust relationship`. 

Use the output of `aws iam-get role` as a starting template.  

* Replace the `Condition.StringEquals` with `Condition.StringLike`.
```shell
"Condition": {
    "StringLike": {
```

* Replace the body of the new `Condition.StringLike` with the provider string.
First, get the `OIDC issuer` with the command `aws eks decribe-cluster --name
<cluster-name>`. 

For example, 
```shell
$ aws eks describe-cluster --name upbound-docs --query "cluster.identity.oidc.issuer" --output text | sed -E 's_^https?://__'
oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4
```
Use this value to build the new contents, in the form:  
`"<oidc issuer>:sub": "system:serviceaccount:<Universal Crossplane
namespace>:provider-aws-*"`

For example, using the previous output:

`"oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4:sub":
"system:serviceaccount:upbound-system:provider-aws-*"`

The value `provider-aws-*` defines the AWS provider and version that needs to
authenticate. Using `*` allows any AWS provider version to authenticate.

An example of the final JSON file.
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Federated": "arn:aws:iam::000000000000:oidc-provider/oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4"
            },
            "Action": "sts:AssumeRoleWithWebIdentity",
            "Condition": {
                "StringLike": {
                    "oidc.eks.us-east-2.amazonaws.com/id/266A01FA1DBF8083FA1C23EB7D4736E4:sub": "system:serviceaccount:upbound-system:provider-aws-*"
                }
            }
        }
    ]
}
```

* Apply the new trust policy. Use the command `aws iam
update-assume-role-policy` to apply the trust policy.

```shell
aws iam update-assume-role-policy \
--role-name <IAM role name> \
--policy-document file://<JSON file>
```

For example,
```shell
aws iam update-assume-role-policy \
--role-name eks-test-role \
--policy-document file://role.json
```

##### Create a ControllerConfig
A `ControllerConfig` creates settings used by the `Provider` deployment.

For IRSA, the `ControllerConfig` provides an `annotation` of the ARN of the role
used by the Kubernetes service account.

First, use `kubectl describe service-account <name> -n upbound-system` to get
the ARN value.

```yaml
$ kubectl describe service-account \
my-upbound-sa \
-n upbound-system
Name:                my-upbound-sa
Namespace:           upbound-system
Labels:              app.kubernetes.io/managed-by=eksctl
Annotations:         eks.amazonaws.com/role-arn: arn:aws:iam::000000000000:role/eks-test-role
Image pull secrets:  <none>
Mountable secrets:   my-upbound-sa-token-spq5k
Tokens:              my-upbound-sa-token-spq5k
Events:              <none>
```

The `Annotations` value is the input for the `ControllerConfig`.

_Note:_ the `ControllerConfig` required for IRSA configuration doesn't require a
`spec` body.

```yaml
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: irsa-controllerconfig
  annotations:
    eks.amazonaws.com/role-arn: arn:aws:iam::000000000000:role/eks-test-role
spec:
```

Apply the `ControllerConfig` with `kubectl apply -f` and verify the installation
with `kubectl get controllerconfig`.

```shell
$ kubectl apply -f controller-config.yml
$ kubectl get controllerconfig
NAME                    AGE
irsa-controllerconfig   6s
```

##### Create a Provider
The `Provider` object references the `ControllerConfig` to use the AWS IAM role
ARN.

The `Provider.spec.controllerConfigRef.name` must match the
`ControllerConfig.name` value. 

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws-s3
spec:
  package: xpkg.upbound.io/upbound/provider-aws-s3:latest
  controllerConfigRef:
    name: irsa-controllerconfig
```

Apply the `Provider` object with `kubectl apply -f` and verify with `kubectl get
providers`.

```shell
$ kubectl apply -f provider.yaml
$ kubectl get providers
NAME              INSTALLED   HEALTHY   PACKAGE                                          AGE
provider-aws-s3   True        True      xpkg.upbound.io/upbound/provider-aws-s3:latest   83s
```

_Note_: it may take up to five minutes for the provider `HEALTHY` value to be
`True`.

##### Create a ProviderConfig
The `ProviderConfig` explicitly configures the official AWS provider-family to use
`IRSA` authentication. 

Define the `ProviderConfig.spec.credentials.source` as `IRSA`. 

```yaml
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: IRSA
```

_Note:_ the value `IRSA` is case sensitive.

Apply the `ProviderConfig` with `kubectl apply -f` and verify with `kubectl get
providerconfigs`.

```shell
$ kubectl apply -f providerconfig.yaml
$ kubectl get providerconfigs
NAME                                    AGE
providerconfig.aws.upbound.io/default   46s
```

The official AWS provider-family now uses the `IRSA` role for authentication to AWS.
