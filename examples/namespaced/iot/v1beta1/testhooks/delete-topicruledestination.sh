#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Note(mbbush): AWS IoT uses the IAM role assigned to the TopicRuleDestination during deletion
# of the TopicRuleDestination to remove the ENIs and other resources it created in the VPC.
# Deletion of the TopicRuleDestination takes several minutes, and if the role is deleted before
# it finishes, then it doesn't have permission to remove the ENIs, so the VPC deletion fails
# because there are still resources left in it. This ordered deletion requirement could be
# encoded in a crossplane Usage resource, but that's still alpha and not readily available
# in the current uptest config. This also solves the problem.
${KUBECTL} delete topicruledestination.iot.aws.upbound.io/iot-topic-rule-destination-example
