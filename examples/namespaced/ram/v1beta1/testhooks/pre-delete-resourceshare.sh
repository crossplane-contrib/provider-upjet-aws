#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(turkenf): We are getting "ResourceShare arn....  could not be found" error  
# when deleted before the ResourceShare resource ResourceShareAccepter resource. 
# This is a workaround for this problem to make automated tests to work.
${KUBECTL} delete resourceshareaccepter.ram.aws.upbound.io -l testing.upbound.io/example-name=example

${KUBECTL} delete principalassociation.ram.aws.upbound.io -l testing.upbound.io/example-name=example