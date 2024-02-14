#!/usr/bin/env bash
set -aeuo pipefail

# Delete the shansot permissions resource before deleting the snapshot istself
${KUBECTL} delete snapshotcreatevolumepermission.ec2.aws.upbound.io --all