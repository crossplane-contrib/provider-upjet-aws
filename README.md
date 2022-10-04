# Upbound AWS Provider

`provider-aws` is a [Crossplane](https://crossplane.io/) provider that is
built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for
[Amazon AWS](https://aws.amazon.com/).

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://github.com/upbound/official-providers/provider-aws/releases):
```
kubectl crossplane install provider upbound/provider-aws:v0.2.1
```

You can see the API reference [here](https://doc.crds.dev/github.com/upbound/official-providers/provider-aws).

## Contributing

Please see the [Adding New Resources](/docs/adding-resources.md) guide.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upbound/official-providers/provider-aws/issues).

## Contact

Please use the following to reach members of the community:

* Slack: Join our [slack channel](https://slack.crossplane.io)
* Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
* Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
* Email: [info@crossplane.io](mailto:info@crossplane.io)

## Governance and Owners

provider-aws is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-aws adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-aws is under the Apache 2.0 license.
