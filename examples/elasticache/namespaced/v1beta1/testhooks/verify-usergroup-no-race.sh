#!/usr/bin/env bash
set -euo pipefail

# SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: CC0-1.0

# After UserGroupAssociation is Ready, force a UserGroup reconcile and verify
# membership is observed without being owned in forProvider (no race with association).
#
# Note: uptest does not inject NAMESPACE into post-assert hooks. Do not honor an
# ambient NAMESPACE from the caller shell — it often points at the wrong ns.

KUBECTL="${KUBECTL:-kubectl}"
NS="upbound-system"
ASSOC_NAME="${ASSOC_NAME:-example}"
GROUP_NAME="${GROUP_NAME:-example}"
USER_ID="${USER_ID:-exampleuserid}"
TIMEOUT_SEC="${TIMEOUT_SEC:-600}"

echo "using namespace ${NS}"
echo "confirming UserGroupAssociation ${ASSOC_NAME} is Ready/Synced"
${KUBECTL} -n "${NS}" wait --for=condition=Ready "usergroupassociation.elasticache.aws.m.upbound.io/${ASSOC_NAME}" --timeout=10m
${KUBECTL} -n "${NS}" wait --for=condition=Synced "usergroupassociation.elasticache.aws.m.upbound.io/${ASSOC_NAME}" --timeout=10m

echo "forcing UserGroup ${GROUP_NAME} reconcile"
${KUBECTL} -n "${NS}" annotate "usergroup.elasticache.aws.m.upbound.io/${GROUP_NAME}" \
  "uptest.upbound.io/force-reconcile=$(date +%s)" --overwrite

echo "asserting forProvider.userIds is unset"
FOR_PROVIDER_USER_IDS="$(${KUBECTL} -n "${NS}" get "usergroup.elasticache.aws.m.upbound.io/${GROUP_NAME}" -o jsonpath='{.spec.forProvider.userIds}')"
if [[ -n "${FOR_PROVIDER_USER_IDS}" ]]; then
  echo "expected spec.forProvider.userIds to be empty, got: ${FOR_PROVIDER_USER_IDS}"
  exit 1
fi

FOR_PROVIDER_REFS="$(${KUBECTL} -n "${NS}" get "usergroup.elasticache.aws.m.upbound.io/${GROUP_NAME}" -o jsonpath='{.spec.forProvider.userIdRefs}')"
if [[ -n "${FOR_PROVIDER_REFS}" ]]; then
  echo "expected spec.forProvider.userIdRefs to be empty, got: ${FOR_PROVIDER_REFS}"
  exit 1
fi

# Ready/Synced may already be True; poll until Observe refreshes membership from AWS.
echo "waiting for status.atProvider.userIds to contain default and ${USER_ID}"
deadline=$((SECONDS + TIMEOUT_SEC))
AT_PROVIDER_USER_IDS=""
while true; do
  AT_PROVIDER_USER_IDS="$(${KUBECTL} -n "${NS}" get "usergroup.elasticache.aws.m.upbound.io/${GROUP_NAME}" -o jsonpath='{.status.atProvider.userIds[*]}')"
  if echo "${AT_PROVIDER_USER_IDS}" | grep -qw default && echo "${AT_PROVIDER_USER_IDS}" | grep -qw "${USER_ID}"; then
    break
  fi
  if (( SECONDS >= deadline )); then
    echo "timeout waiting for atProvider.userIds, got: ${AT_PROVIDER_USER_IDS}"
    exit 1
  fi
  # Retrigger observe periodically (default poll can be up to 10m).
  ${KUBECTL} -n "${NS}" annotate "usergroup.elasticache.aws.m.upbound.io/${GROUP_NAME}" \
    "uptest.upbound.io/force-reconcile=$(date +%s)" --overwrite >/dev/null
  sleep 10
done
echo "atProvider.userIds: ${AT_PROVIDER_USER_IDS}"

echo "asserting UserGroupAssociation ${ASSOC_NAME} is still Ready/Synced"
${KUBECTL} -n "${NS}" wait --for=condition=Ready "usergroupassociation.elasticache.aws.m.upbound.io/${ASSOC_NAME}" --timeout=2m
${KUBECTL} -n "${NS}" wait --for=condition=Synced "usergroupassociation.elasticache.aws.m.upbound.io/${ASSOC_NAME}" --timeout=2m

echo "no-race verification passed"
