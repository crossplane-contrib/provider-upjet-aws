# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Setup k8s tools

-include ../../../build/makelib/k8s_tools.mk

# ====================================================================================
# Setup xpkg

-include ../../../build/makelib/xpkg.mk

# ====================================================================================
# Targets

img.build:
	@$(INFO) Family base image to build: $(IMAGE)
	@$(INFO) Building image $${IMAGE}; \
	$(MAKE) BUILD_ARGS="--load ${BUILD_ARGS}" IMAGE=$${IMAGE} XPKG_REG_ORGS=$(XPKG_REG_ORGS) img.build.shared; \
	if [[ "$${LOAD_PACKAGES}" == "true" ]]; then \
		$(MAKE) batch-process SUBPACKAGES="$(SUBPACKAGES)" BATCH_PLATFORMS=$(PLATFORM) BUILD_ONLY=true STORE_PACKAGES="$$(tr ' ' ',' <<< "$(SUBPACKAGES)")" && \
		for s in $(SUBPACKAGES); do \
			export t=$$(docker load -qi "$(XPKG_OUTPUT_DIR)/$(PLATFORM)/$(PROJECT_NAME)-$${s}-$(VERSION).xpkg") && \
			docker tag $${t##*:} $(BUILD_REGISTRY)/$(PROJECT_NAME)-$${s}-$(ARCH); \
		done; \
	fi || $(FAIL)
	@$(OK) docker build $${IMAGE}

img.build.shared:
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp terraformrc.hcl $(IMAGE_TEMP_DIR) || $(FAIL)
	@docker buildx build $(BUILD_ARGS) \
		--platform $(IMAGE_PLATFORMS) \
		--build-arg TERRAFORM_VERSION=$(TERRAFORM_VERSION) \
		--build-arg TERRAFORM_PROVIDER_SOURCE=$(TERRAFORM_PROVIDER_SOURCE) \
		--build-arg TERRAFORM_PROVIDER_VERSION=$(TERRAFORM_PROVIDER_VERSION) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_NAME=$(TERRAFORM_PROVIDER_DOWNLOAD_NAME) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX=$(TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX) \
		--build-arg CROSSPLANE_PROVIDER_VERSION=$(VERSION) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)

img.publish: batch-process

img.promote:
	@$(INFO) Skipping image promotion from $(FROM_IMAGE) to $(TO_IMAGE)
	@echo Promote is deferred to xpkg machinery
	@$(OK) Image promotion skipped for $(FROM_IMAGE) to $(TO_IMAGE)

CONCURRENCY ?= 30
DEP_CONSTRAINT := >= 0.0.0
ifeq (-,$(findstring -,$(VERSION)))
    DEP_CONSTRAINT = >= 0.0.0-0
endif
BUILD_ONLY ?= false
STORE_PACKAGES ?= ""
batch-process: $(UP)
	@$(INFO) Batch processing smaller provider packages for: "$(SUBPACKAGES)"
	@mkdir -p "$(XPKG_OUTPUT_DIR)/$(PLATFORM)" && \
	$(UP) xpkg batch --smaller-providers "$$(tr ' ' ',' <<< "$(SUBPACKAGES)")" \
		--family-base-image $(BUILD_REGISTRY)/$(PROJECT_NAME) \
		--platform $(BATCH_PLATFORMS) \
		--provider-name $(PROJECT_NAME) \
		--family-package-url-format $(XPKG_REG_ORGS)/%s:$(VERSION) \
		--package-repo-override monolith=$(PROJECT_NAME) --package-repo-override config=provider-family-$(PROVIDER_NAME) \
		--provider-bin-root $(OUTPUT_DIR)/bin \
		--output-dir $(XPKG_OUTPUT_DIR) \
		--store-packages "$(STORE_PACKAGES)" \
		--build-only=$(BUILD_ONLY) \
		--examples-root $(ROOT_DIR)/examples \
		--examples-group-override monolith=* --examples-group-override config=providerconfig \
		--auth-ext $(ROOT_DIR)/package/auth.yaml \
		--crd-root $(ROOT_DIR)/package/crds \
		--crd-group-override monolith=* --crd-group-override config=$(PROVIDER_NAME) \
		--package-metadata-template $(ROOT_DIR)/package/crossplane.yaml.tmpl \
		--template-var XpkgRegOrg=$(XPKG_REG_ORGS) --template-var DepConstraint="$(DEP_CONSTRAINT)" --template-var ProviderName=$(PROVIDER_NAME) \
		--concurrency $(CONCURRENCY) \
		--push-retry 10 || $(FAIL)
	@$(OK) Done processing smaller provider packages for: "$(SUBPACKAGES)"
