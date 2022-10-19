---
title: Quickstart
weight: 1
---
# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and
install the AWS official provider. For more details about the AWS official
provider read the
[Configuration](https://marketplace.upbound.io/providers/upbound/provider-aws/v0.18.0/docs/configuration)
.

To install and use this official provider:
* Install Upbound Universal Crossplane (UXP) into your Kubernetes cluster. 
* Install the `Provider` and apply a `ProviderConfig`.
* Create a *managed resource* in AWS with Kubernetes.

You can walk through this quickstart in one of two ways:
* copy and paste - A list of commands to run to create a managed resource in
  AWS. You can then inspect the Kubernetes cluster and AWS console for more
  information.
* guided tour - A step-by-step walk through of the required commands and
  descriptions on what the commands do.

## Prerequisites
This quickstart requires:
* a Kubernetes cluster with permissions to create pods and secrets
* a host with `kubectl` installed and configured to access the Kubernetes
  cluster
* an AWS account with permissions to create an S3 storage bucket
* AWS [access
  keys](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-creds)

## Copy and paste quickstart

You can either run a single Bash script or run each command individually.

_Note:_ all commands use the current `kubeconfig` context and configuration.

### Bash script
Run the following to download and install 
```shell
curl -sL "https://raw.githubusercontent.com/upbound/provider-aws/v0.18.0/docs/quickstart.sh" | sh
```

### Shell commands

_Note:_ run each command individually or copy to a local to prevent issues
running the commands in the terminal.

```shell
#!/usr/bin/env bash
set -eE

read -p "AWS access_key_id: " aws_access_key; read -sp "AWS secret_access_key: " aws_secret_key; export AWS_KEY=$aws_access_key; export AWS_SECRET=$aws_secret_key; printf "\n"

if ! up --version > /dev/null 2>&1; then printf "Installing up CLI...\n"; curl -sL "https://cli.upbound.io" | sh; sudo mv up /usr/local/bin/; fi

if ! kubectl -n upbound-system get deployment crossplane > /dev/null 2>&1; then printf "Installing UXP...\n" && up uxp install; fi

printf "Checking the UXP installation (this only takes a minute)...\n"
kubectl -n upbound-system wait deployment crossplane --for=condition=Available --timeout=180s


printf "Installing the provider (this will take a few minutes)...\n"
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws
spec:
  package: xpkg.upbound.io/upbound/provider-aws:v0.18.0
EOF
kubectl wait "providers.pkg.crossplane.io/provider-aws" --for=condition=Installed --timeout=180s
kubectl wait "providers.pkg.crossplane.io/provider-aws" --for=condition=Healthy --timeout=180s



cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: aws-secret
  namespace: upbound-system
stringData:
  creds: |
    $(printf "[default]\n    aws_access_key_id = %s\n    aws_secret_access_key = %s" "${AWS_KEY}" "${AWS_SECRET}")
EOF

cat <<EOF | kubectl apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: upbound-system
      name: aws-secret
      key: creds
EOF

cat <<EOF | kubectl create -f -
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  generateName: upbound-bucket-
spec:
  forProvider:
    region: us-east-1
EOF

printf "Checking AWS bucket creation (this only takes a minute)...\n"
kubectl wait "$(kubectl get buckets -o name)" --for=condition=Ready --timeout=180s

kubectl get buckets
```

Your Kubernetes cluster created this AWS S3 bucket.

Remove it with the following command.

```shell
$ kubectl delete bucket --all
```

## Guided tour
These steps are the same as the preceding quickstart, but provides more
information for each action.

_Note:_ all commands use the current `kubeconfig` context and configuration. 

### Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
sudo mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.13.0
```

_Note_: official providers only support `up` command-line versions v0.13.0 or
later.

More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

### Install Upbound Universal Crossplane
Install Upbound Universal Crossplane (UXP) with the Up command-line `up uxp
install` command.

```shell
$ up uxp install
UXP 1.9.0-up.3 installed
```

Verify all UXP pods are `Running` with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                        READY   STATUS    RESTARTS      AGE
crossplane-7fdfbd897c-pmrml                 1/1     Running   0             68m
crossplane-rbac-manager-7d6867bc4d-v7wpb    1/1     Running   0             68m
upbound-bootstrapper-5f47977d54-t8kvk       1/1     Running   0             68m
xgql-7c4b74c458-5bf2q                       1/1     Running   3 (67m ago)   68m
```

_Note:_ `RESTARTS` for the `xgql` pod are normal during initial installation. 

Find more information in the [Upbound UXP
documentation](https://docs.upbound.io/uxp/).

### Install the official AWS provider

Install the official provider into the Kubernetes cluster with a Kubernetes
configuration file. 

```yaml
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws
spec:
  package: xpkg.upbound.io/upbound/provider-aws:v0.18.0
EOF
```

Verify the provider installed with `kubectl describe providers` and `kubectl get
providers`. This `kubectl describe providers` output is from an installed
provider.
```
$ kubectl describe provider

Name:         provider-aws
Namespace:
Labels:       <none>
Annotations:  <none>
API Version:  pkg.crossplane.io/v1
Kind:         Provider
# Output truncated
Status:
  Conditions:
    Last Transition Time:  2022-09-02T20:46:26Z
    Reason:                HealthyPackageRevision
    Status:                True
    Type:                  Healthy
    Last Transition Time:  2022-09-02T20:46:09Z
    Reason:                ActivePackageRevision
    Status:                True
    Type:                  Installed
  Current Identifier:      xpkg.upbound.io/upbound/provider-aws:v0.18.0
  Current Revision:        provider-aws-ab4a3525fb0b
Events:
  Type     Reason                  Age               From                                 Message
  ----     ------                  ----              ----                                 -------
  Warning  InstallPackageRevision  9s (x5 over 14s)  packages/provider.pkg.crossplane.io  current package revision health is unknown
  Normal   InstallPackageRevision  5s (x2 over 5s)    packages/provider.pkg.crossplane.io  Successfully installed package revision
```

The `INSTALLED` value should be `True`. It may take up to 5 minutes for
`HEALTHY` to report true.
```shell
$ kubectl get provider
NAME           INSTALLED   HEALTHY   PACKAGE                                        AGE
provider-aws   True        True   xpkg.upbound.io/upbound/provider-aws:v0.18.0      5s
```

If there are issues downloading and installing the provider the `INSTALLED`
field is empty.

```shell
$ kubectl get providers
NAME           INSTALLED   HEALTHY   PACKAGE                                       AGE
provider-aws                         xpkg.upbound.io/upbound/provider-aws:v0.18.0   62s
```

Use `kubectl describe providers` for more information.

### Create a Kubernetes secret for AWS
The provider requires credentials to create and manage AWS resources.

#### Generate an AWS key-pair file
Create a text file containing the AWS account `aws_access_key_id` and
`aws_secret_access_key`. The [AWS
documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-creds)
provides information on how to generate these keys.

```ini
[default]
aws_access_key_id = <aws_access_key>
aws_secret_access_key = <aws_secret_key>
```

Save this text file as `aws-credentials.txt`.

#### Create a Kubernetes secret with AWS credentials
Use `kubectl create secret -n upbound-system` to generate a Kubernetes secret
object inside the Kubernetes cluster.

```shell
kubectl create secret \
generic aws-secret \
-n upbound-system \
--from-file=creds=./aws-credentials.txt
```

View the secret with `kubectl describe secret`
```shell
$ kubectl describe secret aws-secret -n upbound-system
Name:         aws-secret
Namespace:    upbound-system
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
creds:  114 bytes
```
_Note:_ the size may be larger if there are extra blank space in your text file.

### Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach the AWS
credentials to the installed official provider.

```yaml
cat <<EOF | kubectl apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: upbound-system
      name: aws-secret
      key: creds
EOF
```

The `spec.secretRef` describes the parameters of the secret to use. 
* `namespace` is the Kubernetes namespace the secret is in.
* `name` is the name of the Kubernetes `secret` object.
* `key` is the `Data` field from `kubectl describe secret`.

Apply this configuration with `kubectl apply -f`.

Verify the `ProviderConfig` with `kubectl describe providerconfigs`. 

```yaml
$ kubectl describe providerconfigs
Name:         default
Namespace:
API Version:  aws.upbound.io/v1beta1
Kind:         ProviderConfig
# Output truncated
Spec:
  Credentials:
    Secret Ref:
      Key:        creds
      Name:       aws-secret
      Namespace:  upbound-system
    Source:       Secret
```

**Note:** the `ProviderConfig` install fails and Kubernetes returns an error if
the `Provider` isn't installed.

```shell
$ kubectl apply -f providerconfig.yml
error: resource mapping not found for name: "default" namespace: "" from "providerconfig.yml": no matches for kind "ProviderConfig" in version "aws.upbound.io/v1beta1"
ensure CRDs are installed first
```

### Create a managed resource
Create a managed resource to verify the provider is functioning. 

This example creates an AWS S3 storage bucket, which requires a globally unique
name. 

```yaml
bucket=$(echo "upbound-bucket-"$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 10))
CAT <<EOF | kubectl apply -f -
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  name: $bucket
spec:
  forProvider:
    region: us-east-1
  providerConfigRef:
    name: default
EOF
```

Use `kubectl get buckets` to verify bucket creation.

```shell
$ kubectl get buckets
NAME                           READY   SYNCED   EXTERNAL-NAME                  AGE
upbound-bucket-fb8360b455dd9   True    True     upbound-bucket-fb8360b455dd9   8s
```

Upbound created the bucket when the values `READY` and `SYNCED` are `True`. This
may take up to 5 minutes.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to
understand why.

Here is an example of a failure because the `spec.providerConfigRef.name` value
in the `Bucket` doesn't match the `ProviderConfig` `metadata.name`.

```shell
$ kubectl describe bucket
Name:         upbound-bucket-fb8360b455dd9
Namespace:
Labels:       <none>
Annotations:  crossplane.io/external-name: upbound-bucket-fb8360b455dd9
API Version:  s3.aws.upbound.io/v1beta1
Kind:         Bucket
# Output truncated
Spec:
  Deletion Policy:  Delete
  For Provider:
    Region:  us-east-1
    Tags:
      Crossplane - Kind:            bucket.s3.aws.upbound.io
      Crossplane - Name:            upbound-bucket-fb8360b455dd9
      Crossplane - Providerconfig:  default
  Provider Config Ref:
    Name:  default
Status:
  At Provider:
  Conditions:
    Last Transition Time:  2022-07-25T15:55:41Z
    Message:               connect failed: cannot get terraform setup: cannot get AWS config: cannot get referenced Provider: ProviderConfig.aws.upbound.io "default" not found
    Reason:                ReconcileError
    Status:                False
    Type:                  Synced
Events:
  Type     Reason                   Age              From                                            Message
  ----     ------                   ----             ----                                            -------
  Warning  CannotConnectToProvider  1s (x3 over 2s)  managed/s3.aws.upbound.io/v1beta1, kind=bucket  cannot get terraform setup: cannot get AWS config: cannot get referenced Provider: ProviderConfig.aws.upbound.io "default" not found
```
The output indicates the `Bucket` is using a `ProviderConfig` named `default`.
The applied `ProviderConfig` is `my-config`. 

```shell
$ kubectl get providerconfig
NAME        AGE
providerconfig.aws.upbound.io/my-config   114s
```

### Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `Bucket`
object file. Verify removal of the bucket with `kubectl get buckets`

```shell
$ kubectl delete -f bucket.yml
bucket.s3.aws.upbound.io "upbound-bucket-fb8360b455dd9" deleted

$ kubectl get buckets
No resources found
```
