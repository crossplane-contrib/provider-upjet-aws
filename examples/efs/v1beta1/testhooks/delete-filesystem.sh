#!/usr/bin/env bash
set -aeuo pipefail

# Delete the filesystem resource before deleting the resource itself
${KUBECTL} delete filesystem.efs.aws.upbound.io --all