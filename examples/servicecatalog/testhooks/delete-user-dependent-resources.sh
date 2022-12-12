#!/usr/bin/env bash
set -aeuo pipefail

# Delete the principalportfolioassociation resource before deleting the portfolio
${KUBECTL} delete principalportfolioassociation.servicecatalog.aws.upbound.io --all