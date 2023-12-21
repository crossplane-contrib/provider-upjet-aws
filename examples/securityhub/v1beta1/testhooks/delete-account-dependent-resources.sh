#!/usr/bin/env bash
set -aeuo pipefail

# Delete the findingaggregator, actiontarget, insight resources before deleting the account itself
${KUBECTL} delete findingaggregator.securityhub.aws.upbound.io --all &&
${KUBECTL} delete actiontarget.securityhub.aws.upbound.io --all &&
${KUBECTL} delete insight.securityhub.aws.upbound.io --all