#!/usr/bin/env bash
read -p "Upbound username: " up_user; read -sp "Upbound password: " up_pass; echo ""; read -p "AWS access_key_id: " aws_access_key; read -sp "AWS secret_access_key: " aws_secret_key; export AWS_KEY=$aws_access_key; export AWS_SECRET=$aws_secret_key; export UP_USER=$up_user; export UP_PASS=$up_pass; echo "";

VAR=$(head -n 4096 /dev/urandom | openssl sha1 | tail -c 14)

curl -sL "https://cli.upbound.io" | sh

sudo mv up /usr/local/bin/

up uxp install

printf "\n\nChecking UXP install (this only takes a minute)..." ; while true ; do if [[ $(kubectl get deployment -n upbound-system -o jsonpath='{.items[*].status.conditions[*].status}') = "True True True True True True True True" ]]; then printf "\nUXP is ready.\n\n" ; break; else printf  "."; sleep 2; fi done

up login -u $UP_USER -p $UP_PASS


if [[ $(up org list | wc -l) = 2 ]]; then ORG=$(up org list | awk '{ print $1 }' | sed '1d') ; else up org create my-org-$VAR ; ORG=my-org-$VAR; fi

up robot create my-robot-$VAR -a $ORG

up robot token create my-robot-$VAR my-token-$VAR --output=token.json -a $ORG

up controlplane pull-secret create package-pull-secret -f token.json

cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-aws
spec:
  package: xpkg.upbound.io/upbound/provider-aws:v0.10.0
  packagePullSecrets:
    - name: package-pull-secret
EOF

printf "\n\nChecking provider (this will take a few minutes)..." ; while true ; do if [[ $(kubectl get provider -o jsonpath='{.items[*].status.conditions[*].status}') = "True True" ]]; then printf "\n\n The provider is ready.\n\n" ; break; else echo  -n "."; sleep 3; fi done

cat <<EOF > aws.txt
[default]
aws_access_key_id = $AWS_KEY
aws_secret_access_key = $AWS_SECRET
EOF

kubectl create secret generic aws-secret -n upbound-system --from-file=creds=./aws.txt

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

cat <<EOF | kubectl apply -f -
apiVersion: s3.aws.upbound.io/v1beta1
kind: Bucket
metadata:
  name: upbound-bucket-$VAR
spec:
  forProvider:
    region: us-east-1
  providerConfigRef:
    name: default
EOF

printf "\n\nChecking AWS bucket creation (this only takes a minute)..." ; while true ; do if [[ $(kubectl get buckets -o jsonpath='{.items[*].status.conditions[*].status}') = "True True" ]]; then printf "\nYour bucket is ready.\n\n" ; break; else printf  "."; sleep 2; fi done

kubectl get buckets