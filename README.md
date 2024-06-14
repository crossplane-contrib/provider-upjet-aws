<!--
SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>

SPDX-License-Identifier: CC-BY-4.0
-->

This branch is created to trigger Uptest. Do not merge!

# Official AWS Provider

<div align="center">

![CI](https://github.com/crossplane-contrib/provider-upjet-aws/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/crossplane-contrib/provider-upjet-aws/all.svg)](https://github.com/crossplane-contrib/provider-upjet-aws/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/crossplane-contrib/provider-upjet-aws)](https://goreportcard.com/report/github.com/crossplane-contrib/provider-upjet-aws)
[![Contributors](https://img.shields.io/github/contributors/crossplane-contrib/provider-upjet-aws)](https://github.com/crossplane-contrib/provider-upjet-aws/graphs/contributors)
[![Slack](https://img.shields.io/badge/Slack-4A154B?logo=slack)](https://crossplane.slack.com/archives/C05E0UE46S2)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/crossplane_io)](https://twitter.com/crossplane_io)

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

### Add a New Resource

Follow the guide [here](https://github.com/crossplane/upjet/blob/v0.10.0/docs/add-new-resource-short.md).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-aws/issues/new/choose).

## Contact

[#upjet-provider-aws](https://crossplane.slack.com/archives/C05E0UE46S2) channel in
[Crossplane Slack](https://slack.crossplane.io)

## Licensing

Provider AWS is under [the Apache 2.0 license](LICENSE) with [notice](NOTICE).
