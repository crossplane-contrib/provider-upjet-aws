#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete the principalportfolioassociation resource before deleting the portfolio
${KUBECTL} delete principalportfolioassociation.servicecatalog.aws.upbound.io --all