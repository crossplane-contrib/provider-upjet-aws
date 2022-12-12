#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

if [[ -n "${UPTEST_CLOUD_CREDENTIALS:-}" ]]; then
  # UPTEST_CLOUD_CREDENTIALS may contain more than one cloud credentials that we expect to be provided
  # in a single GitHub secret. We expect them provided as key=value pairs separated by newlines. Currently we expect
  # two AWS IAM user credentials to be provided. For example:
  # DEFAULT='[default]
  # aws_access_key_id = REDACTED
  # aws_secret_access_key = REDACTED'
  # PEER='[default]
  # aws_access_key_id = REDACTED
  # aws_secret_access_key = REDACTED'
  eval "${UPTEST_CLOUD_CREDENTIALS}"

  if [[ -n "${DEFAULT:-}" ]]; then
    echo "Creating the default cloud credentials secret..."
    ${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="${DEFAULT}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

    echo "Creating a default provider config..."
    cat <<EOF | ${KUBECTL} apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
EOF
  fi

  if [[ -n "${PEER:-}" ]]; then
    echo "Creating the peer cloud credentials secret for cross-account testing..."
    ${KUBECTL} -n upbound-system create secret generic provider-secret-peer --from-literal=credentials="${PEER}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

    echo "Creating a peer provider config for cross-account testing..."
    cat <<EOF | ${KUBECTL} apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: peer
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret-peer
      namespace: upbound-system
      key: credentials
EOF
  fi
fi
