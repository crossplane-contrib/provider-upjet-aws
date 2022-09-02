# Quickstart

This guide walks through the process to install Upbound Universal Crossplane and install the AWS official provider. For more details about the AWS official provider read the [Configuration](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/docs/configuration) .

To install and use this official provider:
* Create an upbound.io account.
* Install the `up` command-line.
* Install Upbound Universal Crossplane (UXP) into your Kubernetes cluster. 
* Authenticate to the Upbound Marketplace and generate a Kubernetes secret.
* Install the `Provider` and apply a `ProviderConfig`.
* Create a *managed resource* in AWS with Kubernetes.

## Create an Upbound.io user account
Create an account on [Upbound.io](https://accounts.upbound.io/register). 

## Install the Up command-line
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

_Note_: official providers only support `up` command-line versions v0.13.0 or later.

More information about the Up command-line is available in the [Upbound Up documentation](https://docs.upbound.io/cli/).

## Install Upbound Universal Crossplane
Install Upbound Universal Crossplane (UXP) with the Up command-line.

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

Find more information in the [Upbound UXP documentation](https://docs.upbound.io/uxp/).

## Log in with the Up command-line
Use `up login` to authenticate to the Upbound Marketplace.

It's important to use `-a <your organization>` when logging in. Only accounts belonging to organizations can use official providers.

```shell
$ up login -a my-org
username: my-user
password: 
my-user logged in
```

## Create an Upbound robot account
Upbound robots are identities used for authentication that are independent from a single user and aren't tied to specific usernames or passwords.

Creating a robot account allows Kubernetes to install an official provider.

Create a new robot account with the command
```shell
up robot create \
<robot account name>
```

_Note_: only users logged into an organization can create robot accounts.

```shell
$ up robot create my-robot
my-org/my-robot created
```

## Create an Upbound robot account token
The token associates with a specific robot account and acts as a username and password for authentication.

Generate a token using the command
```shell
up robot token create \
<robot account> \
<token name> \
--output=<file>
```

```shell
$ up robot token create my-robot my-token --output=token.json
my-org/my-robot/my-token created
```

The `output` file is a JSON file containing the robot token's `accessId` and `token`. The `accessId` is the username and `token` is the password for the token.

_Note_: you can't recover a lost robot token. You must delete and recreate the token.


## Create a Kubernetes pull secret
Downloading and installing official providers requires Kubernetes to authenticate to the Upbound Marketplace using a Kubernetes `secret` object.

Use the `up` command-line to generate a Kubernetes secret using your robot account and token. 
```shell
up controlplane \
pull-secret create \
package-pull-secret \
-f <robot token file> 
```

Provide a name for your Kubernetes secret and the robot token JSON file.

For example,
```shell
$ up controlplane pull-secret create package-pull-secret -f token.json
my-org/package-pull-secret created
```

_Note_: you must provide the robot token file or you can't authenticate to install an official provider.  

`up` creates the secret in the `upbound-system` namespace. 

```shell
$ kubectl get secret -n upbound-system
NAME                                         TYPE                             DATA   AGE
package-pull-secret                            kubernetes.io/dockerconfigjson   1      8m46s
sh.helm.release.v1.universal-crossplane.v1   helm.sh/release.v1               1      21m
upbound-agent-tls                            Opaque                           3      21m
uxp-ca                                       Opaque                           3      21m
xgql-tls                                     Opaque                           3      21m
```

## Install the official AWS provider
Install the official provider into the Kubernetes cluster with a Kubernetes configuration file. 

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws
spec:
  package: xpkg.upbound.io/upbound/provider-aws:latest
  packagePullSecrets:
    - name: package-pull-secret
```

_Note_: the `name` of the `packagePullSecrets` must be the same as the name of the Kubernetes secret just created.

Apply this configuration with `kubectl apply -f`.

After installing the provider, verify the install with `kubectl get providers`.   

```shell
$ kubectl get providers
NAME           INSTALLED   HEALTHY   PACKAGE                                       AGE
provider-aws   True        True      xpkg.upbound.io/upbound/provider-aws:latest   62s
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
Use `kubectl create secret -n upbound-system` to generate a Kubernetes secret object inside the Kubernetes cluster.

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

**Note:** the `ProviderConfig` install fails and Kubernetes returns an error if the `Provider` isn't installed, for example, due to incorrect credentials provided via `packagePullSecrets`.

```shell
$ kubectl apply -f providerconfig.yml
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

Use `kubectl get buckets` to verify bucket creation.

```shell
$ kubectl get buckets
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
The output indicates the `Bucket` is using a `ProviderConfig` named `default`. The applied `ProviderConfig` is `my-config`. 

```shell
$ kubectl get providerconfig
NAME        AGE
providerconfig.aws.upbound.io/my-config   114s
```

## Delete the managed resource
Remove the managed resource by using `kubectl delete -f` with the same `Bucket` object file. Verify removal of the bucket with `kubectl get buckets`

```shell
$ kubectl delete -f bucket.yml
bucket.s3.aws.upbound.io "upbound-bucket-fb8360b455dd9" deleted

$ kubectl get buckets
No resources found
```
