Install official provider on self-hosted control planes using Universal Crossplane (`UXP`).

## Install the provider
Official providers require a Kubernetes `imagePullSecret` to install. 
<!-- vale gitlab.Substitutions = NO --> 
<!-- Details on creating an `imagePullSecret` are available in the [generic provider documentation](/providers/#create-a-kubernetes-imagepullsecret) -->
<!-- vale gitlab.Substitutions = YES --> 

_Note:_ if you already installed an official provider using an `imagePullSecret` a new secret isn't required.

Install the Upbound official AWS provider with the following configuration file

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws
spec:
  package: xpkg.upbound.io/upbound/provider-aws:v0.5.0
  packagePullSecrets:
    - name: package-pull-secret
```

Define the provider version with `spec.package`. This example uses version `v0.5.0`.

The `spec.packagePullSecrets.name` value matches the Kubernetes `imagePullSecret`. The secret must be in the same namespace as the Upbound pod.

Install the provider with `kubectl apply -f`.

Verify the configuration with `kubectl get provider`.

```shell
$ kubectl get providers
NAME           INSTALLED   HEALTHY   PACKAGE                                       AGE
provider-aws   True        True      xpkg.upbound.io/upbound/provider-aws:v0.5.0   62s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/crossplane/crossplane/pkg.crossplane.io/Provider/v1) to view all available `Provider` options.

## Configure the provider
The AWS provider requires credentials for authentication to AWS. The AWS provider consumes the credentials from a Kubernetes secret object.

### Generate a Kubernetes secret
To create the Kubernetes secret, create a text file containing the AWS account `aws_access_key_id` and `aws_secret_access_key`. The [AWS documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-creds)

```ini
[default]
aws_access_key_id = <aws_access_key>
aws_secret_access_key = <aws_secret_key>
```

Create the secret with the command  

`kubectl create secret generic <secret name> --from-file=<aws_credentials_file.txt>`

### Create a ProviderConfig object
Apply the secret in a `ProviderConfig` Kubernetes configuration file.

```yaml
apiVersion: aws.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
    name: default
spec:
    credentials:
        source: Secret
        secretRef:
            namespace: upbound-system
            name: aws-secret
            key: aws-secret
```

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the `kubectl create secret generic <name>` command.

View the [ProviderConfig CRD definition](resources/aws.upbound.io/ProviderConfig/v1beta1) to view all available `ProviderConfig` options.
