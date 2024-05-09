#!/usr/bin/env bash
set -aeuo pipefail

# Note(mbbush): AWS uses the IAM role to remove the ENIs and other resources it created in the VPC.
# If the kinesisdataanalytics application has been started, it creates ENIs, so it needs to delete them.
# If the role is deleted before the application deletion finishes, then it doesn't have permission
# to remove the ENIs, so the VPC deletion fails because there are still resources left in it.
# This ordered deletion requirement could be encoded in a crossplane Usage resource, but that's
# still alpha and not readily available in the current uptest config. This also solves the problem.
${KUBECTL} delete application.kinesisanalyticsv2.aws.upbound.io/example-application-snapshot
