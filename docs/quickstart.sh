#!/usr/bin/env bash
set -eE

read -p "AWS access_key_id: " aws_access_key; read -sp "AWS secret_access_key: " aws_secret_key; export AWS_KEY=$aws_access_key; export AWS_SECRET=$aws_secret_key; printf "\n"

if ! up --version > /dev/null 2>&1; then printf "Installing up CLI...\n"; curl -sL "https://cli.upbound.io" | sh; sudo mv up /usr/local/bin/; fi

if ! kubectl -n upbound-system get deployment crossplane > /dev/null 2>&1; then printf "Installing UXP...\n" && up uxp install; fi

printf "Checking the UXP installation (this only takes a minute)...\n"
kubectl -n upbound-system wait deployment crossplane --for=condition=Available --timeout=180s


printf "Installing the provider (this will take a few minutes)...\n"
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws
spec:
  package: xpkg.upbound.io/upbound/provider-aws:v0.17.0
EOF
kubectl wait "providers.pkg.crossplane.io/provider-aws" --for=condition=Installed --timeout=180s
kubectl wait "providers.pkg.crossplane.io/provider-aws" --for=condition=Healthy --timeout=180s

creds=$(
cat <<EOF | base64
[default]
aws_access_key_id = $AWS_KEY
aws_secret_access_key = $AWS_SECRET
EOF
)

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: aws-secret
  namespace: upbound-system
data:
  creds: ${creds}
EOF

cat <<EOF | kubectl apply -f -
apiVersion: aws.upbound.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      namespace: upbound-system
      name: aws-secret
      key: creds
EOF

cat <<EOF | kubectl create -f -
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  generateName: upbound-bucket-
spec:
  forProvider:
    region: us-east-1
EOF

printf "Checking AWS bucket creation (this only takes a minute)...\n"
kubectl wait "$(kubectl get buckets -o name)" --for=condition=Ready --timeout=180s

kubectl get buckets
