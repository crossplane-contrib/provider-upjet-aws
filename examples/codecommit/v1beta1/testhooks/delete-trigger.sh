#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(turkenf): We are getting "RepositoryDoesNotExist" exception if repository for
# the trigger got deleted before the trigger resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete trigger.codecommit.aws.upbound.io/example