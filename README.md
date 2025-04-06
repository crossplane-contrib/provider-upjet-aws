<!--
SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>

SPDX-License-Identifier: CC-BY-4.0
-->

# Upjet-based Crossplane provider for AWS

<div style="text-align: center;">

![CI](https://github.com/crossplane-contrib/provider-upjet-aws/workflows/CI/badge.svg)
[![GitHub release](https://img.shields.io/github/release/crossplane-contrib/provider-upjet-aws/all.svg)](https://github.com/crossplane-contrib/provider-upjet-aws/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/crossplane-contrib/provider-upjet-aws)](https://goreportcard.com/report/github.com/crossplane-contrib/provider-upjet-aws)
[![Contributors](https://img.shields.io/github/contributors/crossplane-contrib/provider-upjet-aws)](https://github.com/crossplane-contrib/provider-upjet-aws/graphs/contributors)
[![Slack](https://img.shields.io/badge/Slack-4A154B?logo=slack)](https://crossplane.slack.com/archives/C05E0UE46S2)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/crossplane_io)](https://twitter.com/crossplane_io)

</div>

Provider Upjet-AWS is a [Crossplane](https://crossplane.io/) provider that is
built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for
[Amazon AWS](https://aws.amazon.com/).

## Getting Started

Follow the quick start
guide [here](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/docs/quickstart).

You can find a detailed API reference for all the managed resources with examples in
the [Upbound Marketplace](https://marketplace.upbound.io/providers/upbound/provider-aws/latest/managed-resources).

For more information about monitoring the Upjet runtime, please
see [Monitoring Guide](https://github.com/crossplane/upjet/blob/main/docs/monitoring.md)

## Contributing

For the general contribution guide,
see [Upjet Contribution Guide](https://github.com/crossplane/upjet/blob/main/CONTRIBUTING.md)

If you'd like to learn how to use Upjet, see [Usage Guide](https://github.com/crossplane/upjet/tree/main/docs).

To build this provider locally and run it in a local Kubernetes cluster, run the
following to build the family config provider (`config`) and the sub-package you are working on e.g. `ec2`:

```shell
DOCKERHUB_ORG=<your-docker-name>
BUILD_ARGS="--load" XPKG_REG_ORGS_NO_PROMOTE="" XPKG_REG_ORGS="index.docker.io/$DOCKERHUB_ORG" make build.all publish BRANCH_NAME=main SUBPACKAGES="config ec2"
```

The `BRANCH_NAME` is set to `main` (even though you might be on another branch) to allow
for the publishing of the images to your docker hub account.

To install the provider, in this example `provider-aws-ec2`, into a local Kubernetes cluster with Crossplane already
installed, apply:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws-ec2
spec:
  package: docker.io/<your-docker-name>/provider-aws-ec2:<the-version-taken-from-the-output-of-the-previous-command>
```

### Add a New Resource

Follow the Upjet guide
for [adding new resources](https://github.com/crossplane/upjet/blob/main/docs/adding-new-resource.md).

### Changing the Upstream Terraform Provider
Provider Upjet-AWS depends on and is generated from terraform-provider-aws. It makes a
few changes to the source of the terraform provider to expose internal functionality,
or fix crossplane-specific issues. These changes are contained in `.patch` files in
the `patches` directory. These patches are automatically applied to the source in the
`upstream` git submodule before running any of the makefile targets, using logic
defined in `scripts/upstream.sh`.

Here is one example workflow for making changes to the upstream provider:
```
# Convert the patches in the patch directory into commits, and check out the 
# detached head created by applying them on top of the commit in the parent repo
./scripts/upstream.sh checkout

# change the code as needed
cd upstream
...

# commit your changes, and make/amend existing commits as needed
git commit --whatever

# convert the git commits back into patches, overwriting what's currently in the
# patches directory
cd ..
./scripts/upstream.sh check_in

# At this point the upstream submodule has the base git ref checked out (the 
# unmodified upstream terraform provider), but with all the crossplane-specific
# patches applied to the source code. The git tree is dirty.

# test your changes
...

# Commit your changes to the patch files to provider-upjet-aws
git commit -s -m "Changed the terraform provider to make it glow in the dark"
```

It is important that the git reference stored in this repo for the `upstream` submodule
is an unmodified ref from the upstream repository. The only time it should change is
when upgrading the upstream provider version.

## Getting help

For filing bugs, suggesting improvements, or requesting new resources or features, please
open an [issue](https://github.com/crossplane-contrib/provider-upjet-aws/issues/new/choose).

For general help on using the provider consider asking the Crossplane community in the
[#upjet-provider-aws](https://crossplane.slack.com/archives/C05E0UE46S2) channel in
[Crossplane Slack](https://slack.crossplane.io)

## License

The provider is released under the [the Apache 2.0 license](LICENSE) with [notice](NOTICE).
