#!/usr/bin/env bash

set -eux

# setting up colors
BLU='\033[0;34m'
YLW='\033[0;33m'
GRN='\033[0;32m'
RED='\033[0;31m'
NOC='\033[0m' # No Color
echo_info(){
    printf "\n${BLU}%s${NOC}" "$1"
}
echo_step(){
    printf "\n${BLU}>>>>>>> %s${NOC}\n" "$1"
}
echo_sub_step(){
    printf "\n${BLU}>>> %s${NOC}\n" "$1"
}

echo_step_completed(){
    printf "${GRN} [✔]${NOC}"
}

echo_success(){
    printf "\n${GRN}%s${NOC}\n" "$1"
}
echo_warn(){
    printf "\n${YLW}%s${NOC}" "$1"
}
echo_error(){
    printf "\n${RED}%s${NOC}" "$1"
    exit 1
}

# ------------------------------
projectdir="$( cd "$( dirname "${BASH_SOURCE[0]}")"/.. && pwd )"

# get the build environment variables from the special build.vars target in the main makefile
eval $(make --no-print-directory -C ${projectdir} build.vars)

# ------------------------------

CROSSPLANE_NAMESPACE="crossplane-system"
K8S_CLUSTER="automated-tests"

# setup package cache
echo_step "setting up local package cache"
CACHE_PATH="${projectdir}/.work/automated-tests-package-cache"
mkdir -p "${CACHE_PATH}"
echo "created cache dir at ${CACHE_PATH}"

"${UP}" xpkg xp-extract --from-xpkg ./"${PROVIDER_NAME}"/_output/xpkg/"${PLATFORM}"/"${PROVIDER_NAME}"-*.xpkg -o "${CACHE_PATH}/${PROVIDER_NAME}.gz" && chmod 644 "${CACHE_PATH}/${PROVIDER_NAME}.gz"

# create kind cluster with extra mounts
KIND_NODE_IMAGE="kindest/node:${KIND_NODE_IMAGE_TAG}"
echo_step "creating k8s cluster using kind and node image ${KIND_NODE_IMAGE}"
KIND_CONFIG="$( cat <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraMounts:
  - hostPath: "${CACHE_PATH}/"
    containerPath: /cache
EOF
)"
echo "${KIND_CONFIG}" | "${KIND}" create cluster --name="${K8S_CLUSTER}" --wait=5m --image="${KIND_NODE_IMAGE}" --config=-

# tag controller image and load it into kind cluster
cd ./"${PROVIDER_NAME}"
BUILD_REGISTRY="build-$(echo "${HOSTNAME}"-"$(pwd)" | shasum -a 256 | cut -c1-8)"
cd "${projectdir}"

BUILD_IMAGE="${BUILD_REGISTRY}/${PROVIDER_NAME}-${SAFEHOSTARCH}"
PACKAGE_IMAGE="${PROVIDER_NAME}:${VERSION}"

docker tag "${BUILD_IMAGE}" "${PACKAGE_IMAGE}"
"${KIND}" load docker-image "${PACKAGE_IMAGE}" --name="${K8S_CLUSTER}"

echo_step "create crossplane-system namespace"
"${KUBECTL}" create ns crossplane-system

echo_step "create persistent volume and claim for mounting package-cache"
PV_YAML="$( cat <<EOF
apiVersion: v1
kind: PersistentVolume
metadata:
  name: package-cache
  labels:
    type: local
spec:
  storageClassName: manual
  capacity:
    storage: 5Mi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/cache"
EOF
)"
echo "${PV_YAML}" | "${KUBECTL}" create -f -

PVC_YAML="$( cat <<EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: package-cache
  namespace: crossplane-system
spec:
  accessModes:
    - ReadWriteOnce
  volumeName: package-cache
  storageClassName: manual
  resources:
    requests:
      storage: 1Mi
EOF
)"
echo "${PVC_YAML}" | "${KUBECTL}" create -f -

# install crossplane from stable channel
# TODO(hasheddan): switch to using UXP for all testing
echo_step "installing crossplane from stable channel"
"${HELM3}" repo add crossplane-stable https://charts.crossplane.io/stable/
chart_version="$("${HELM3}" search repo crossplane-stable/crossplane | awk 'FNR == 2 {print $2}')"
echo_info "using crossplane version ${chart_version}"
echo
# we replace empty dir with our PVC so that the /cache dir in the kind node
# container is exposed to the crossplane pod
"${HELM3}" install crossplane --namespace crossplane-system crossplane-stable/crossplane --version ${chart_version} --wait --set packageCache.pvc=package-cache

# ----------- integration tests
echo_step "--- INTEGRATION TESTS ---"

# install package
echo_step "installing ${PROVIDER_NAME} into \"${CROSSPLANE_NAMESPACE}\" namespace"

CONFIG_YAML="$( cat <<EOF
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: config
spec:
  image: "${PACKAGE_IMAGE}"
  args: ["-d"]
  env:
  - name: UPBOUND_CONTEXT
    value: testing
EOF
)"

INSTALL_YAML="$( cat <<EOF
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: "${PROVIDER_NAME}"
spec:
  package: "${PROVIDER_NAME}"
  packagePullPolicy: Never
  controllerConfigRef:
    name: config
EOF
)"

echo "${CONFIG_YAML}" | "${KUBECTL}" apply -f -
echo "${INSTALL_YAML}" | "${KUBECTL}" apply -f -

# printing the cache dir contents can be useful for troubleshooting failures
echo_step "check kind node cache dir contents"
docker exec "${K8S_CLUSTER}-control-plane" ls -la /cache

echo_step "waiting for provider to be installed"

kubectl wait "provider.pkg.crossplane.io/${PROVIDER_NAME}" --for=condition=healthy --timeout=180s

kubectl get deployment -n crossplane-system
