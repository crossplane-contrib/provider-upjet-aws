#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Delete the shansot permissions resource before deleting the snapshot istself
${KUBECTL} delete snapshotcreatevolumepermission.ec2.aws.upbound.io --all