---
title: Quickstart
weight: 1
---

# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and 
install the AWS official provider-family.

To use AWS official provider-family, install Upbound Universal Crossplane into your 
Kubernetes cluster, install the `Provider`, apply a `ProviderConfig`, and create 
a *managed resource* in AWS via Kubernetes.

## Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
sudo mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.19.1
```

_Note_: official providers only support `up` command-line versions v0.13.0 or
later.

More information about the Up command-line is available in the [Upbound Up
documentation](https://docs.upbound.io/cli/).

## Install Upbound Universal Crossplane
Install Upbound Universal Crossplane (UXP) with the Up command-line `up uxp
install` command.

```shell
$ up uxp install
UXP 1.13.2-up.2 installed
```

Verify all UXP pods are `Running` with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                       READY   STATUS    RESTARTS   AGE
crossplane-77ff754998-4l8xb                1/1     Running   0          21s
crossplane-rbac-manager-79b8bdd6d8-ml6ft   1/1     Running   0          21s
```

Find more information in the [Upbound UXP
documentation](https://docs.upbound.io/uxp/).

## Install the official AWS provider-family

Install the official provider-family into the Kubernetes cluster with a Kubernetes
configuration file. For instance, let's install the `provider-aws-s3`

_Note_: The first provider installed of a family also installs an extra provider-family Provider.
The provider-family provider manages the ProviderConfig for all other providers in the same family.

```yaml
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws-s3
spec:
  package: xpkg.upbound.io/upbound/provider-aws-s3:<version>
EOF
```

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
NAME                          INSTALLED   HEALTHY   PACKAGE                                               AGE
provider-aws-s3               True        True      xpkg.upbound.io/upbound/provider-aws-s3:v0.41.0       6m39s
upbound-provider-family-aws   True        True      xpkg.upbound.io/upbound/provider-family-aws:v0.41.0   6m30s
```

It may take up to 5 minutes to report `HEALTHY`.

If you are going to use your own registry please check [Install Providers in an offline environment](https://docs.upbound.io/providers/provider-families/#installing-a-provider-family:~:text=services%20to%20install.-,Install%20Providers%20in%20an%20offline%20environment,-View%20the%20installed)

## Create a Kubernetes secret for AWS
The official provider-family requires credentials to create and manage AWS resources.

### Generate an AWS key-pair file
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

### Create a Kubernetes secret with AWS credentials
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

## Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach the AWS
credentials to the installed official `provider-aws-s3`.

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

## Create a managed resource
Create a managed resource to verify the `provider-aws-s3` is functioning. 

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

## Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `Bucket`
object file. Verify removal of the bucket with `kubectl get buckets`

```shell
$ kubectl delete -f bucket.yml
bucket.s3.aws.upbound.io "upbound-bucket-fb8360b455dd9" deleted

$ kubectl get buckets
No resources found
```
