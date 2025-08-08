#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(turkenf): We are getting "GraphQL API owci22ox5vgivmcrcmoe7khjla not found" exception if graphqlapi for
# the apikey got deleted before the apikey resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete apikey.appsync.aws.upbound.io/example