#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete the findingaggregator, actiontarget, insight resources before deleting the account itself
${KUBECTL} delete findingaggregator.securityhub.aws.upbound.io --all &&
${KUBECTL} delete actiontarget.securityhub.aws.upbound.io --all &&
${KUBECTL} delete insight.securityhub.aws.upbound.io --all