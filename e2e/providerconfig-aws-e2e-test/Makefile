# Project Setup
PROJECT_NAME := providerconfig-aws-e2e-test
PROJECT_REPO := github.com/upbound/provider-upjet-aws

# NOTE(hasheddan): the platform is insignificant here as Configuration package
# images are not architecture-specific. We constrain to one platform to avoid
# needlessly pushing a multi-arch image.
PLATFORMS ?= linux_amd64
-include ../../build/makelib/common.mk

# ====================================================================================
# Setup Kubernetes tools

KIND_VERSION = v0.22.0
UP_VERSION = v0.28.0
UP_CHANNEL = stable
UPTEST_VERSION = v0.11.1
YQ_VERSION = v4.40.5

-include ../../build/makelib/k8s_tools.mk
# ====================================================================================
# Setup XPKG
XPKG_DIR = $(shell pwd)/package
XPKG_IGNORE = .github/workflows/*.yaml,.github/workflows/*.yml,examples/*.yaml,.work/uptest-datasource.yaml,.cache/**,_output/**
XPKG_REG_ORGS ?= xpkg.upbound.io/upbound
# NOTE(hasheddan): skip promoting on xpkg.upbound.io as channel tags are
# inferred.
XPKG_REG_ORGS_NO_PROMOTE ?= xpkg.upbound.io/upbound
XPKGS = $(PROJECT_NAME)
-include ../../build/makelib/xpkg.mk

CROSSPLANE_NAMESPACE = upbound-system
CROSSPLANE_ARGS = "--enable-usages,--debug"
-include ../../build/makelib/local.xpkg.mk
-include ../../build/makelib/controlplane.mk


# ====================================================================================
# Targets

# run `make help` to see the targets and options

# We want submodules to be set up the first time `make` is run.
# We manage the build/ folder and its Makefiles as a submodule.
# The first time `make` is run, the includes of build/*.mk files will
# all fail, and this target will be run. The next time, the default as defined
# by the includes will be run instead.
fallthrough: submodules
	@echo Initial setup complete. Running make again . . .
	@make

# Update the submodules, such as the common build scripts.
submodules:
	@git submodule sync
	@git submodule update --init --recursive

# We must ensure up is installed in tool cache prior to build as including the k8s_tools machinery prior to the xpkg
# machinery sets UP to point to tool cache.
build.init: $(UP)

# ====================================================================================
# End to End Testing

# This target requires the following environment variables to be set:
# - UPTEST_CLOUD_CREDENTIALS, cloud credentials for the provider being tested, e.g. export UPTEST_CLOUD_CREDENTIALS=$(cat ~/.aws/credentials)
# - To ensure the proper functioning of the end-to-end test resource pre-deletion hook, it is crucial to arrange your resources appropriately.
#   You can check the basic implementation here: https://github.com/upbound/uptest/blob/main/internal/templates/01-delete.yaml.tmpl.
# - UPTEST_DATASOURCE_PATH (optional), see https://github.com/upbound/uptest#injecting-dynamic-values-and-datasource
uptest: $(UPTEST) $(KUBECTL) $(KUTTL) $(YQ)
	@$(INFO) running automated tests
	@KUBECTL=$(KUBECTL) KUTTL=$(KUTTL) CROSSPLANE_NAMESPACE=$(CROSSPLANE_NAMESPACE) $(UPTEST) e2e package/examples/e2etestcluster-claim.yaml --data-source="${UPTEST_DATASOURCE_PATH}" --setup-script=test/setup.sh --default-timeout=5400 || $(FAIL)
	@$(OK) running automated tests

# This target requires the following environment variables to be set:
# - UPTEST_CLOUD_CREDENTIALS, cloud credentials for the provider being tested, e.g. export UPTEST_CLOUD_CREDENTIALS=$(cat ~/.aws/credentials)
e2e: build controlplane.up local.xpkg.deploy.configuration.$(PROJECT_NAME) uptest-e2e

e2e-lite: build controlplane.up local.xpkg.deploy.configuration.$(PROJECT_NAME)

uptest-e2e: $(UPTEST) $(KUBECTL) $(KUTTL) $(YQ)
	@$(INFO) dump e2e claim:
	@mkdir -p "_output"
	@$(YQ) '(.spec.parameters.targetClusterParameters.provider.familyPackage = env(AWS_FAMILY_PACKAGE_IMAGE)) | \
		   (.spec.parameters.targetClusterParameters.provider.ec2Package = env(AWS_EC2_PACKAGE_IMAGE)) | \
		   (.spec.parameters.targetClusterParameters.provider.rdsPackage = env(AWS_RDS_PACKAGE_IMAGE)) | \
		   (.spec.parameters.targetClusterParameters.crossplaneVersion = env(TARGET_CROSSPLANE_VERSION)) ' \
				package/examples/e2etestcluster-claim.yaml > '_output/e2etestcluster-claim.yaml'
	if [ -n "${AWS_EKS_IAM_DEFAULT_ADMIN_ROLE}" ]; \
	then \
  		echo "overriding EKS cluster default IAM role from environment";\
  		$(YQ) -i '(.spec.parameters.iam.roleArn = env(AWS_EKS_IAM_DEFAULT_ADMIN_ROLE)) ' _output/e2etestcluster-claim.yaml;\
  	fi
	@cat _output/e2etestcluster-claim.yaml
	@$(INFO) running automated tests
	@KUBECTL=$(KUBECTL) KUTTL=$(KUTTL) CROSSPLANE_NAMESPACE=$(CROSSPLANE_NAMESPACE) $(UPTEST) e2e _output/e2etestcluster-claim.yaml --data-source="${UPTEST_DATASOURCE_PATH}" --setup-script=test/setup.sh --default-timeout=5400 || $(FAIL)
	@$(OK) running automated tests

render:
	crossplane beta render package/examples/e2etestcluster-claim.yaml package/apis/e2etestcluster/composition.yaml package/examples/functions.yaml -r

yamllint:
	@$(INFO) running yamllint
	@yamllint ./apis || $(FAIL)
	@$(OK) running yamllint

.PHONY: uptest e2e render yamllint
