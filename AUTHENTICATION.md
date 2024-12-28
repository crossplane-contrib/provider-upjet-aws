<!--
SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>

SPDX-License-Identifier: CC-BY-4.0
-->

## Authentication mechanisms supported by provider-aws

`provider-aws` needs to authenticate itself to the AWS API. Like other
Crossplane providers, the credentials to be used during authentication can be
configured by means of `aws.upbound.io/v1beta1/ProviderConfig` resources.
`provider-aws` currently supports the following authentication mechanisms:
- Authentication with long-term IAM user credentials
- Authentication using *I*AM *R*oles for *S*ervice *A*ccounts (IRSA)
- Authentication using EKS Pod Identity
- Authentication using an assumed Web identity

The authentication mechanism to be used can be selected by setting the
`spec.credentials.source` field of the `ProviderConfig` to one of the following
values:
- `Secret`
- `IRSA`
- `WebIdentity`
to configure long-term credentials, IRSA and authentication with an assumed Web
identity, respectively.
- `PodIdentity`
to configure authentication by utilizing EKS Pod Identity to assume an IAM role.

If no authentication mechanism is specified, the default is to use the
`Secret` authentication mechanism.

*Note*: Please note that with Upbound managed control-planes (MCP), we will only support
the `Secret` authentication mechanism with `provider-aws` until EKS clusters can
be utilized.

Each authentication mechanism may need further configuration specific to it, and
the configuration details of each mechanism will be detailed below together with some
example `ProviderConfig` manifests. And orthogonal to the selected
authentication mechanism's configuration, one can also configure AWS IAM role
chaining [[1]] to assume a series of roles and use the short-term credentials of
those assumed roles. A detailed discussion on role chaining and the
authentication & authorization scenarios enabled by it and shortcomings of it
are out of the scope of this document but a detailed account can be found in
[[2]].

In the following sections, we will first discuss the configuration options for
the currently supported authentication mechanisms for the provider, and then
describe how IAM role chaining can be configured with each authentication
method.

### Authentication with long-term IAM user credentials
This authentication method involves making an IAM user's long-term credentials
available to `provider-aws` by means of a Kubernetes `Secret` and thus is
discouraged from a security perspective. Details can be found in [[3]]. The
required configuration is a reference to a Kubernetes secret containing the
long-term credentials. And example `Secret` configuration is as follows:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-aws-creds
  namespace: upbound-system
type: Opaque
data:
  credentials: <base64-encoded credentials file>
---
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-aws-creds
      namespace: upbound-system
      key: credentials
```
The base64-encoded credentials document should be formatted basically as a
shared AWS config file [[4]] with the `[default]` profile:

```
[default]
aws_access_key_id=<AWS_ACCESS_KEY_ID>
aws_secret_access_key=<AWS_SECRET_ACCESS_KEY>
```

As mentioned above, this authentication method involves IAM user's long-term
credentials and is considered as less secure when compared to other
configurations.

### Authentication using IRSA
IRSA authentication is available when `provider-aws` is running on an EKS
cluster and [IRSA has been configured for that
cluster](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html).
Configuring IRSA for EKS involves associating a Kubernetes `ServiceAccount` with
an IAM role so that an EKS workload running under that `ServiceAccount` will be
authenticated as its associated IAM Role against the AWS API. The association
between the Kubernetes `ServiceAccount` and the IAM role is done by annotating
the `ServiceAccount` with a `eks.amazonaws.com/role-arn` key. An example is as
follows:

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    eks.amazonaws.com/role-arn: arn:aws:iam::111122223333:role/iam-role-name
```

Thus, we need to have the `ServiceAccount` under which `provider-aws` is running
annotated with the key `eks.amazonaws.com/role-arn`, and this can be accomplished
in a number of ways. One approach is to use a `DeploymentRuntimeConfig` while
installing the provider:

```yaml
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: irsa-drc
spec:
  serviceAccountTemplate:
    metadata:
      annotations:
        eks.amazonaws.com/role-arn: arn:aws:iam::111122223333:role/iam-role-name
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: upbound-provider-aws-ec2
spec:
  package: xpkg.upbound.io/upbound/provider-aws-ec2:v1.4.0
  runtimeConfigRef:
    name: irsa-drc
```

After the `ServiceAccount` under which `provider-aws` is running is annotated,
the following `ProviderConfig` can be used to utilize IRSA authentication for
the provider:

```yaml
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: irsa
spec:
  credentials:
    source: IRSA
```

As it can be seen in the example `ProviderConfig` above, no explicit credentials
need to be referred by the `ProviderConfig` and EKS kubelets running on the
nodes are responsible for injecting & renewing short-term credentials needed to
authenticate the provider as the associated IAM Role. This authentication
method, thus, is considered as more secure when compared to the `Secret` method
discussed above. IRSA requires a trust relationship to be defined between the
EKS cluster's OIDC provider and the associated IAM Role (whose ARN is specified
in the `ServiceAccount`'s annotation). Please refer to [[5]] for detailed
configuration instructions.

### Web Identity
It’s now possible to specify the WebIdentity tokens to be used in `ProviderConfig`s
for WebIdentity authentication. Before `v1.1.0`, it was only possible to do so via
the environment variables.

`ProviderConfig` API specification is expanded with `spec.credentials.webIdentity.tokenConfig`
which allows consumers to configure the token to be used for WebIdentity authentication.
Consumers can reference a secret or filesystem location for the token to be used for
`WebIdentity` authentication.

- Each `ProviderConfig` using WebIdentity authentication can now use different tokens
per `ProviderConfig` object, allowing multiple WebIdentity configurations in a single cluster.

- ℹ️ The change is backward compatible for consumers relying on the old behavior where they
set both of the `AWS_WEB_IDENTITY_TOKEN_FILE` and `AWS_ROLE_ARN` [[6]] environment variables. When
`spec.credentials.webIdentity.tokenConfig` is not specified, the old behavior is assumed.

- ⚠️ Deprecation Notice: Configuring the WebIdentity authentication using the
`AWS_WEB_IDENTITY_TOKEN_FILE` and `AWS_ROLE_ARN` environment variables is now deprecated
in favor of the new `spec.credentials.webIdentity.tokenConfig` API.

An example WebIdentity token configuration where the token is read from a Kubernetes secret
is as follows:

```yaml
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: webidentity-example
spec:
  credentials:
    source: WebIdentity
    webIdentity:
      roleARN: arn:aws:iam::123456789012:role/providerexamplerole
      tokenConfig:
        source: Secret
        secretRef:
          key: token
          name: example-web-identity-token-secret
          namespace: upbound-system
```

Another example using a filesystem location is as follows:
```yaml
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: webidentity-example
spec:
  credentials:
    source: WebIdentity
    webIdentity:
      roleARN: arn:aws:iam::123456789012:role/providerexamplerole
      tokenConfig:
        source: Filesystem
        fs:
          path: /path/to/token/file
```

Please note that the `Filesystem` source option needs the token to be mounted as a file
in the filesystem of the provider pod, e.g,. via a `DeploymentRuntimeConfig`.

The difference is that the new API effectively allows specifying the token per `ProviderConfig`.

### IAM Role chaining configuration
With all of the authentication mechanisms discussed above, `provider-aws` allows
IAM role chaining to be configured [[1]]. Omitting the actual authentication
mechanism configuration, a `ProviderConfig` with IAM role chaining configuration
looks like the following:

```yaml
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: config-with-role-chaining
spec:
  credentials:
    source: ...
  assumeRoleChain:
    - roleARN: arn:aws:iam::111111111111:role/account-1-iam-role-name
    - roleARN: arn:aws:iam::222222222222:role/account-2-iam-role-name
```

For a detailed account on the tenant isolation strategies enabled with role
chaining and especially with using Web Identity method in tandem with role
chaining, please refer to [[2]]. But with the above configuration, the provider
first assumes (*not* considering a potential initial
`sts.AssumeRoleWithWebIdentity` call) the `account-1-iam-role-name` IAM Role
in the AWS account `111111111111` and then, using `account-1-iam-role-name`'s
temporary credentials, assumes `account-2-iam-role-name` in the account
`222222222222`. There must exist appropriate trust relationships between those
accounts to allow role chaining and the final account on this chain
(`arn:aws:iam::222222222222:role/account-2-iam-role-name` in the above example
chain) must have the required privileges on the AWS API to manage the AWS
resources. Please note that this is an ordered list of IAM Roles, which must
match the chain of the trust policies defined among the roles.

### EKS Pod Identity

EKS Pod Identity authentication is available when `provider-aws` is running on an EKS cluster and [EKS Pod Identity has been configured for that cluster](https://docs.aws.amazon.com/eks/latest/userguide/pod-identities.html). Unlike IRSA, EKS Pod Identity eliminates the need for an OIDC provider. Instead, it relies on the built-in `pods.eks.amazonaws.com` service principal and the EKS Pod Identity Agent for managing IAM roles and credentials.

Configuring EKS Pod Identity for EKS involves: 
1) [Installing EKS Pod Identity Agent](https://docs.aws.amazon.com/eks/latest/userguide/pod-id-agent-setup.html)
2) Associating a Kubernetes `ServiceAccount` with an IAM role so that an EKS workload running under that `ServiceAccount` will be authenticated as its associated IAM Role against the AWS API. [The association between the Kubernetes `ServiceAccount` and the IAM role](https://docs.aws.amazon.com/eks/latest/userguide/pod-id-association.html#pod-id-association-create) is done by creating an EKS Pod Identity association between the `ServiceAccount` and `namespace` on the EKS cluster and the IAM role on the AWS account.

The `ServiceAccount` under which `provider-aws` is running, and the `namespace` in which the `provider-aws` is deployed, must match the configuration of the previously configured Pod Identity association. 

It can be done by using a `DeploymentRuntimeConfig` while installing the provider:

```yaml
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
    name: podidentity-drc
spec:
    serviceAccountTemplate:
        metadata:
            name: provider-aws
    deploymentTemplate: {}

---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: upbound-provider-aws-ec2
spec:
  package: xpkg.upbound.io/upbound/provider-aws-ec2:v1.4.0
  runtimeConfigRef:
    name: podidentity-drc
```

`ProviderConfig` can be configured to use EKS Pod Identity authentication by setting the `source` field to `PodIdentity`: 

```yaml
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
    name: podidentity-example
spec:
    credentials:
        source: PodIdentity
```

[EKS Pod Identity troubleshooting tips](https://docs.aws.amazon.com/eks/latest/userguide/addon-id-troubleshoot.html):
- Validate the `namespace` and `ServiceAccount` match the Pod Identity association.
- Confirm that the Pod Identity Agent has been deployed and is functional.
- Review the IAM role's trust policy for Pod Identity

[1]:
    https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html

[2]:
    https://github.com/crossplane/upjet/blob/main/docs/design-doc-provider-identity-based-auth.md

[3]:
    https://marketplace.upbound.io/providers/upbound/provider-aws/v0.37.0/docs/configuration

[4]: https://docs.aws.amazon.com/sdkref/latest/guide/file-format.html

[5]:
    https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html

[6]:
    https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRoleWithWebIdentity.html
