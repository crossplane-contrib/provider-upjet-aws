#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(mbbush): AWS Batch  uses the IAM role assigned to the Compute Environment during deletion
# of the Compute Environment.
# Deletion of the Compute Environment takes several minutes, and if the role is deleted before
# it finishes, then it doesn't have permission, so the Compute Environment deletion fails
# because there are still resources left in it. This ordered deletion requirement could be
# encoded in a crossplane Usage resource, but that's still alpha and not readily available
# in the current uptest config. This also solves the problem.
${KUBECTL} delete computeenvironment.batch.aws.upbound.io/sample