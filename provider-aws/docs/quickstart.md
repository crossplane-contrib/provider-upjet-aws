The Upbound AWS Provider is the officially supported provider for Amazon Web Services (AWS).

View the AWS Provider Documentation configuration for details and configuration options. 

## Quickstart
This guide walks through the process to install Upbound Universal Crossplane and install the AWS official provider.

To use this official provider, install Upbound Universal Crossplane into your Kubernetes cluster, install the `Provider`, apply a `ProviderConfiguration`, and create a *managed resource* in AWS via Kubernetes.

## Create an Upbound.io user account
Create an account on [Upbound.io](https://accounts.upbound.io/register). 

## Install the Up command-line
Download and install the Upbound `up` command-line.

```shell
curl -sL "https://cli.upbound.io" | sh
mv up /usr/local/bin/
```

Verify the version of `up` with `up --version`

```shell
$ up --version
v0.13.0
```

_Note_: official providers only support `up` command-line versions v0.13.0 or later.

## Install Universal Crossplane
Install Upbound Universal Crossplane with the Up command-line.

```shell
$ up uxp install
UXP 1.9.0-up.3 installed
```

Verify the UXP pods are running with `kubectl get pods -n upbound-system`

```shell
$ kubectl get pods -n upbound-system
NAME                                        READY   STATUS    RESTARTS      AGE
crossplane-7fdfbd897c-pmrml                 1/1     Running   0             68m
crossplane-rbac-manager-7d6867bc4d-v7wpb    1/1     Running   0             68m
upbound-bootstrapper-5f47977d54-t8kvk       1/1     Running   0             68m
xgql-7c4b74c458-5bf2q                       1/1     Running   3 (67m ago)   68m
```

## Log in with the Up command-line
Use `up login` to authenticate to the Upbound Marketplace.

It's important to use `-a <your organization>` when logging in. Only accounts belonging to organizations can use official providers.

```shell
$ up login -a my-org
username: my-user
password: 
my-user logged in
```
## Create a Kubernetes pull secret
Downloading and installing official providers requires Kubernetes to authenticate to the Upbound Marketplace using a Kubernetes `secret` object.

Using the `up controlplane pull-secret` command creates an [Upbound robot account](http://docs.upbound.io/cli/command-reference/robot/) account. 

_Note_: robot accounts are independent from your account. Your account information is never stored in Kubernetes.

```shell
$ up controlplane pull-secret create my-upbound-secret
WARNING: Using temporary user credentials that will expire within 30 days.
my-org/my-upbound-secret created
```

`Up` creates the secret in the `upbound-system` namespace. 

```shell
$ kubectl get secret -n upbound-system
NAME                                         TYPE                             DATA   AGE
my-upbound-secret                            kubernetes.io/dockerconfigjson   1      8m46s
sh.helm.release.v1.universal-crossplane.v1   helm.sh/release.v1               1      21m
upbound-agent-tls                            Opaque                           3      21m
uxp-ca                                       Opaque                           3      21m
xgql-tls                                     Opaque                           3      21m
```

## Install the official AWS provider in to the managed control plane
<!-- Use the marketplace button -->

Install the official provider into the managed control plane with a Kubernetes configuration file. 

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws
spec:
  package: xpkg.upbound.io/upbound/provider-aws:v0.8.0
  packagePullSecrets:
    - name: my-upbound-secret
```

_Note_: the `name` of the `packagePullSecrets` must be the same as the name of the Kubernetes secret just created.

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
$ kubectl get providers
NAME           INSTALLED   HEALTHY   PACKAGE                                       AGE
provider-aws   True        True      xpkg.upbound.io/upbound/provider-aws:v0.8.0   62s
```

It may take up to 5 minutes to report `HEALTHY`.

If the `packagePullSecrets` is incorrect the provider returns a `401 Unauthorized` error. View the status and error with `kubectl describe provider`.

```yaml
$ kubectl describe provider
Name:         provider-aws
API Version:  pkg.crossplane.io/v1
Kind:         Provider
# Output truncated
Events:
  Type     Reason         Age              From                                 Message
  ----     ------         ----             ----                                 -------
  Warning  UnpackPackage  1s (x4 over 9s)  packages/provider.pkg.crossplane.io  cannot unpack package: failed to fetch package digest from remote: GET https://xpkg.upbound.io/service/token?scope=repository%!A(MISSING)upbound%!F(MISSING)provider-aws%!A(MISSING)pull&service=xpkg.upbound.io: unexpected status code 401 Unauthorized
```

## Create a Kubernetes secret for AWS
The provider requires credentials to create and manage AWS resources.

### Generate an AWS key-pair file
Create a text file containing the AWS account `aws_access_key_id` and `aws_secret_access_key`. The [AWS documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-creds) provides information on how to generate these keys.

```ini
[default]
aws_access_key_id = <aws_access_key>
aws_secret_access_key = <aws_secret_key>
```

Save this text file as `aws-credentials.txt`.

### Create a Kubernetes secret with AWS credentials
Use `kubectl create secret -n upbound-system` to generate the Kubernetes secret object inside the managed control plane.

`kubectl create secret generic aws-secret -n upbound-system --from-file=creds=./aws-credentials.txt`

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
## Create a ProviderConfig
Create a `ProviderConfig` Kubernetes configuration file to attach the AWS credentials to the installed official provider.

```yaml
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
```

Apply this configuration with `kubectl apply -f`.

**Note:** the `Providerconfig` value `spec.secretRef.name` must match the `name` of the secret in `kubectl get secrets -n upbound-system` and `spec.SecretRef.key` must match the value in the `Data` section of the secret.

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

**Note:** the `ProviderConfig` install fails and Kubernetes returns an error if the `Provider` isn't installed, for example, due to `packagePullSecrets` authentication.

```shell
vagrant@kubecontroller-01:~$ kubectl apply -f providerconfig.yml
error: resource mapping not found for name: "default" namespace: "" from "providerconfig.yml": no matches for kind "ProviderConfig" in version "aws.upbound.io/v1beta1"
ensure CRDs are installed first
```

## Create a managed resource
Create a managed resource to verify the provider is functioning. 

This example creates an AWS S3 storage bucket, which requires a globally unique name. 

Generate a unique bucket name from the command line.

`echo "upbound-bucket-"$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 14)`

For example
```
$ echo "upbound-bucket-"$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 10)
upbound-bucket-fb8360b455dd9
```

Use this bucket name for `metadata.name` value.

Create a `Bucket` configuration file. Replace `<BUCKET NAME>` with the `upbound-bucket-` generated name.

```yaml
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  name: <BUCKET NAME>
spec:
  forProvider:
    region: us-east-1
  providerConfigRef:
    name: default
```

**Note:** the `spec.providerConfigRef.name` must match the `ProviderConfig` `metadata.name` value.

Apply this configuration with `kubectl apply -f`.

Use `kubectl get bucket` to verify bucket creation.

```shell
$ kubectl get bucket
NAME                           READY   SYNCED   EXTERNAL-NAME                  AGE
upbound-bucket-fb8360b455dd9   True    True     upbound-bucket-fb8360b455dd9   8s
```

Upbound created the bucket when the values `READY` and `SYNCED` are `True`.

If the `READY` or `SYNCED` are blank or `False` use `kubectl describe` to understand why.

Here is an example of a failure because the `spec.providerConfigRef.name` value in the `Bucket` doesn't match the `ProviderConfig` `metadata.name`.

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
The output indicates the `Bucket` is using `ProviderConfig` named `default`. The applied `ProviderConfig` is `my-config`. 

```shell
$ kubectl get providerconfig
NAME        AGE
providerconfig.aws.upbound.io/my-config   114s
```

## Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `Bucket` object file. Verify removal of the bucket with `kubectl get bucket`

```shell
$ kubectl get bucket
No resources found
```