# Official AWS Provider

<div align="center">

![CI](https://github.com/upbound/provider-aws/workflows/CI/badge.svg) [![GitHub release](https://img.shields.io/github/release/upbound/provider-aws/all.svg?style=flat-square)](https://github.com/upbound/provider-aws/releases) [![Go Report Card](https://goreportcard.com/badge/github.com/upbound/provider-aws)](https://goreportcard.com/report/github.com/upbound/provider-aws) [![Slack](https://slack.crossplane.io/badge.svg)](https://crossplane.slack.com/archives/C01TRKD4623) [![Twitter Follow](https://img.shields.io/twitter/follow/upbound_io.svg?style=social&label=Follow)](https://twitter.com/intent/follow?screen_name=upbound_io&user_id=788180534543339520)

</div>

Provider AWS is a [Crossplane](https://crossplane.io/) provider that is
built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for
[Amazon AWS](https://aws.amazon.com/).

## Getting Started

Follow the quick start guide [here](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/docs/quickstart).

You can find a detailed API reference for all the managed resources with examples in the [Upbound Marketplace](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/managed-resources).

For getting more information about resource consumption and monitoring
the upjet runtime, please see [Sizing Guide](https://github.com/crossplane/upjet/blob/v0.10.0/docs/sizing-guide.md)
and [Monitoring Guide](https://github.com/crossplane/upjet/blob/main/docs/monitoring.md)

## Contributing

For the general contribution guide, see [Upjet Contribution Guide](https://github.com/crossplane/upjet/blob/main/CONTRIBUTING.md)

If you'd like to learn how to use Upjet, see [Usage Guide](https://github.com/crossplane/upjet/tree/main/docs).

To build this provider locally and run it in a local Kubernetes cluster, run the
following to build the family provider (`config`) and `ec2`:

```shell
DOCKERHUB_ORG=<your-docker-name>
BUILD_ARGS="--load" XPKG_REG_ORGS_NO_PROMOTE="" XPKG_REG_ORGS="index.docker.io/$DOCKERHUB_ORG" make build.all publish BRANCH_NAME=main SUBPACKAGES="config ec2"
```

The `BRANCH_NAME` set to `main` (even though you might be on another branch) will
let make publish the images to your docker hub account.

To install the provider `provider-aws-ec2` into a local Kubernetes cluster, apply:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws-ec2
spec:
  package: docker.io/<your-docker-name>/provider-aws-ec2:<the-version-taken-from-the-output-of-the-previous-command>
```

Use `monolith` instead of `ec2` to build and publish the monolithic provider.

### Add a New Resource

Follow the guide [here](https://github.com/crossplane/upjet/blob/v0.10.0/docs/add-new-resource-short.md).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/upbound/provider-aws/issues).

## Contact

Please open a Github issue for all requests. If you need to reach out to Upbound,
you can do so via the following channels:
* Slack: [#upbound](https://crossplane.slack.com/archives/C01TRKD4623) channel in [Crossplane Slack](https://slack.crossplane.io)
* Twitter: [@upbound_io](https://twitter.com/upbound_io)
* Email: [support@upbound.io](mailto:support@upbound.io)

## Licensing

Provider AWS is under [the Apache 2.0 license](LICENSE) with [notice](NOTICE).
