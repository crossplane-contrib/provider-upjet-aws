#!/usr/bin/env bash
set -aeuo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# Runs the update test as a post-assert hook instead of the
# uptest.upbound.io/update-parameter annotation, whose JSON uptest v2.2.0
# corrupts before passing it to kubectl (crossplane/uptest#66).
mr=secret.secretsmanager.aws.upbound.io/example

${KUBECTL} --subresource=status patch "${mr}" --type=merge -p '{"status":{"conditions":[]}}'
${KUBECTL} patch "${mr}" --type=merge -p '{"spec":{"forProvider":{"tags":{"updated-by":"crossplane"}}}}'
${KUBECTL} wait "${mr}" --for=condition=Test --timeout=10m
${KUBECTL} wait "${mr}" --for=jsonpath='{.status.atProvider.tags.updated-by}'=crossplane --timeout=10m
