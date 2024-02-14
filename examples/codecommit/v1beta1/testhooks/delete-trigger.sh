#!/usr/bin/env bash
set -aeuo pipefail

# Note(turkenf): We are getting "RepositoryDoesNotExist" exception if repository for
# the trigger got deleted before the trigger resource. This is a workaround for this
# problem to make automated tests to work.
${KUBECTL} delete trigger.codecommit.aws.upbound.io/example