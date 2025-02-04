#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Waiting until all configuration packages are healthy/installed..."
"${KUBECTL}" wait configuration.pkg --all --for=condition=Healthy --timeout 5m
"${KUBECTL}" wait configuration.pkg --all --for=condition=Installed --timeout 5m
"${KUBECTL}" wait configurationrevisions.pkg --all --for=condition=Healthy --timeout 5m

echo "Creating cloud credential secret..."
"${KUBECTL}" -n upbound-system create secret generic aws-creds --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" \
    --dry-run=client -o yaml | "${KUBECTL}" apply -f -

echo "Waiting until all installed provider packages are healthy..."
"${KUBECTL}" wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
"${KUBECTL}" -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Waiting for all XRDs to be established..."
"${KUBECTL}" wait xrd --all --for condition=Established

echo "Creating a default provider config..."
cat <<EOF | "${KUBECTL}" apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    secretRef:
      key: credentials
      name: aws-creds
      namespace: upbound-system
    source: Secret
EOF
