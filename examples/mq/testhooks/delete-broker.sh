#!/usr/bin/env bash
set -aeuo pipefail

# Delete the broker resource before deleting the secret
${KUBECTL} delete broker.mq.aws.upbound.io --all