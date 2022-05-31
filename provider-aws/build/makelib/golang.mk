# Copyright 2016 The Upbound Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# ====================================================================================
# Options

# The go project including repo name, for example, github.com/rook/rook
GO_PROJECT ?= $(PROJECT_REPO)

# Optional. These are subdirs that we look for all go files to test, vet, and fmt
GO_SUBDIRS ?= cmd pkg

# Optional. Additional subdirs used for integration or e2e testings
GO_INTEGRATION_TESTS_SUBDIRS ?=

# Optional directories (relative to CURDIR)
GO_VENDOR_DIR ?= vendor
GO_PKG_DIR ?= $(WORK_DIR)/pkg

# Optional build flags passed to go tools
GO_BUILDFLAGS ?=
GO_LDFLAGS ?=
GO_TAGS ?=
GO_TEST_FLAGS ?=
GO_TEST_SUITE ?=
GO_NOCOV ?=
GO_COVER_MODE ?= count
GO_CGO_ENABLED ?= 0

# ====================================================================================
# Setup go environment

# turn on more verbose build when V=1
ifeq ($(V),1)
GO_LDFLAGS += -v -n
GO_BUILDFLAGS += -x
endif

# whether to generate debug information in binaries. this includes DWARF and symbol tables.
ifeq ($(DEBUG),0)
GO_LDFLAGS += -s -w
endif

# set GOOS and GOARCH
GOOS := $(OS)
GOARCH := $(ARCH)
export GOOS GOARCH

# set GOHOSTOS and GOHOSTARCH
GOHOSTOS := $(HOSTOS)
GOHOSTARCH := $(TARGETARCH)

GO_PACKAGES := $(foreach t,$(GO_SUBDIRS),$(GO_PROJECT)/$(t)/...)
GO_INTEGRATION_TEST_PACKAGES := $(foreach t,$(GO_INTEGRATION_TESTS_SUBDIRS),$(GO_PROJECT)/$(t)/integration)

ifneq ($(GO_TEST_PARALLEL),)
GO_TEST_FLAGS += -p $(GO_TEST_PARALLEL)
endif

ifneq ($(GO_TEST_SUITE),)
GO_TEST_FLAGS += -run '$(GO_TEST_SUITE)'
endif

GOPATH := $(shell go env GOPATH)

# setup tools used during the build
DEP_VERSION=v0.5.1
DEP := $(TOOLS_HOST_DIR)/dep-$(DEP_VERSION)
GOJUNIT := $(TOOLS_HOST_DIR)/go-junit-report
GOCOVER_COBERTURA := $(TOOLS_HOST_DIR)/gocover-cobertura
GOIMPORTS := $(TOOLS_HOST_DIR)/goimports

GO := go
GOHOST := GOOS=$(GOHOSTOS) GOARCH=$(GOHOSTARCH) go
GO_VERSION := $(shell $(GO) version | sed -ne 's/[^0-9]*\(\([0-9]\.\)\{0,4\}[0-9][^.]\).*/\1/p')

# we use a consistent version of gofmt even while running different go compilers.
# see https://github.com/golang/go/issues/26397 for more details
GOFMT_VERSION := 1.11
ifneq ($(findstring $(GOFMT_VERSION),$(GO_VERSION)),)
GOFMT := $(shell which gofmt)
else
GOFMT := $(TOOLS_HOST_DIR)/gofmt$(GOFMT_VERSION)
endif

# We use a consistent version of golangci-lint to ensure everyone gets the same
# linters.
GOLANGCILINT_VERSION ?= 1.31.0
GOLANGCILINT := $(TOOLS_HOST_DIR)/golangci-lint-v$(GOLANGCILINT_VERSION)

GO_BIN_DIR := $(abspath $(OUTPUT_DIR)/bin)
GO_OUT_DIR := $(GO_BIN_DIR)/$(PLATFORM)
GO_TEST_DIR := $(abspath $(OUTPUT_DIR)/tests)
GO_TEST_OUTPUT := $(GO_TEST_DIR)/$(PLATFORM)
GO_LINT_DIR := $(abspath $(OUTPUT_DIR)/lint)
GO_LINT_OUTPUT := $(GO_LINT_DIR)/$(PLATFORM)

ifeq ($(GOOS),windows)
GO_OUT_EXT := .exe
endif

ifeq ($(RUNNING_IN_CI),true)
# Output checkstyle XML rather than human readable output.
# the timeout is increased to 10m, to accommodate CI machines with low resources.
GO_LINT_ARGS := --timeout 10m0s --out-format=checkstyle > $(GO_LINT_OUTPUT)/checkstyle.xml
endif

# NOTE: the install suffixes are matched with the build container to speed up the
# the build. Please keep them in sync.

# we run go build with -i which on most system's would want to install packages
# into the system's root dir. using our own pkg dir avoid thats
ifneq ($(GO_PKG_DIR),)
GO_PKG_BASE_DIR := $(abspath $(GO_PKG_DIR)/$(PLATFORM))
GO_PKG_STATIC_FLAGS := -pkgdir $(GO_PKG_BASE_DIR)_static
endif

GO_COMMON_FLAGS = $(GO_BUILDFLAGS) -tags '$(GO_TAGS)'
GO_STATIC_FLAGS = $(GO_COMMON_FLAGS) $(GO_PKG_STATIC_FLAGS) -installsuffix static  -ldflags '$(GO_LDFLAGS)'
GO_GENERATE_FLAGS = $(GO_BUILDFLAGS) -tags 'generate $(GO_TAGS)'

export GO111MODULE

# switch for go modules
ifeq ($(GO111MODULE),on)

# set GOPATH to $(GO_PKG_DIR), so that the go modules are installed there, instead of the default $HOME/go/pkg/mod
export GOPATH=$(abspath $(GO_PKG_DIR))

endif

# ====================================================================================
# Go Targets

go.init: go.vendor.lite
	@if [ "$(GO111MODULE)" != "on" ] && [ "$(realpath ../../../..)" !=  "$(realpath $(GOPATH))" ]; then \
		$(WARN) the source directory is not relative to the GOPATH at $(GOPATH) or you are you using symlinks. The build might run into issue. Please move the source directory to be at $(GOPATH)/src/$(GO_PROJECT) ;\
	fi

go.build:
	@$(INFO) go build $(PLATFORM)
	$(foreach p,$(GO_STATIC_PACKAGES),@CGO_ENABLED=0 $(GO) build -v -i -o $(GO_OUT_DIR)/$(lastword $(subst /, ,$(p)))$(GO_OUT_EXT) $(GO_STATIC_FLAGS) $(p) || $(FAIL) ${\n})
	$(foreach p,$(GO_TEST_PACKAGES) $(GO_LONGHAUL_TEST_PACKAGES),@CGO_ENABLED=0 $(GO) test -i -c -o $(GO_TEST_OUTPUT)/$(lastword $(subst /, ,$(p)))$(GO_OUT_EXT) $(GO_STATIC_FLAGS) $(p) || $(FAIL) ${\n})
	@$(OK) go build $(PLATFORM)

go.install:
	@$(INFO) go install $(PLATFORM)
	$(foreach p,$(GO_STATIC_PACKAGES),@CGO_ENABLED=0 $(GO) install -v $(GO_STATIC_FLAGS) $(p) || $(FAIL) ${\n})
	@$(OK) go install $(PLATFORM)

go.test.unit: $(GOJUNIT) $(GOCOVER_COBERTURA)
	@$(INFO) go test unit-tests
ifeq ($(GO_NOCOV),true)
	@$(WARN) coverage analysis is disabled
	@CGO_ENABLED=0 $(GOHOST) test $(GO_TEST_FLAGS) $(GO_STATIC_FLAGS) $(GO_PACKAGES) || $(FAIL)
else
	@mkdir -p $(GO_TEST_OUTPUT)
	@CGO_ENABLED=$(GO_CGO_ENABLED) $(GOHOST) test -i -cover $(GO_STATIC_FLAGS) $(GO_PACKAGES) || $(FAIL)
	@CGO_ENABLED=$(GO_CGO_ENABLED) $(GOHOST) test -v -covermode=$(GO_COVER_MODE) -coverprofile=$(GO_TEST_OUTPUT)/coverage.txt $(GO_TEST_FLAGS) $(GO_STATIC_FLAGS) $(GO_PACKAGES) 2>&1 | tee $(GO_TEST_OUTPUT)/unit-tests.log || $(FAIL)
	@cat $(GO_TEST_OUTPUT)/unit-tests.log | $(GOJUNIT) -set-exit-code > $(GO_TEST_OUTPUT)/unit-tests.xml || $(FAIL)
	@$(GOCOVER_COBERTURA) < $(GO_TEST_OUTPUT)/coverage.txt > $(GO_TEST_OUTPUT)/coverage.xml
endif
	@$(OK) go test unit-tests

# Depends on go.test.unit, but is only run in CI with a valid token after unit-testing is complete
# DO NOT run locally.
go.test.codecov:
	@$(INFO) go test codecov
	@cd $(GO_TEST_OUTPUT) && bash <(curl -s https://codecov.io/bash) || $(FAIL)
	@$(OK) go test codecov

go.test.integration: $(GOJUNIT)
	@$(INFO) go test integration-tests
	@mkdir -p $(GO_TEST_OUTPUT) || $(FAIL)
	@CGO_ENABLED=0 $(GOHOST) test -i $(GO_STATIC_FLAGS) $(GO_INTEGRATION_TEST_PACKAGES) || $(FAIL)
	@CGO_ENABLED=0 $(GOHOST) test $(GO_TEST_FLAGS) $(GO_STATIC_FLAGS) $(GO_INTEGRATION_TEST_PACKAGES) $(TEST_FILTER_PARAM) 2>&1 | tee $(GO_TEST_OUTPUT)/integration-tests.log || $(FAIL)
	@cat $(GO_TEST_OUTPUT)/integration-tests.log | $(GOJUNIT) -set-exit-code > $(GO_TEST_OUTPUT)/integration-tests.xml || $(FAIL)
	@$(OK) go test integration-tests

go.lint: $(GOLANGCILINT)
	@$(INFO) golangci-lint
	@mkdir -p $(GO_LINT_OUTPUT)
	@$(GOLANGCILINT) run $(GO_LINT_ARGS) || $(FAIL)
	@$(OK) golangci-lint

go.vet:
	@$(INFO) go vet $(PLATFORM)
	@CGO_ENABLED=0 $(GOHOST) vet $(GO_COMMON_FLAGS) $(GO_PACKAGES) $(GO_INTEGRATION_TEST_PACKAGES) || $(FAIL)
	@$(OK) go vet $(PLATFORM)

go.fmt: $(GOFMT)
	@$(INFO) go fmt
	@gofmt_out=$$($(GOFMT) -s -d -e $(GO_SUBDIRS) $(GO_INTEGRATION_TESTS_SUBDIRS) 2>&1) && [ -z "$${gofmt_out}" ] || (echo "$${gofmt_out}" 1>&2; $(FAIL))
	@$(OK) go fmt

go.fmt.simplify: $(GOFMT)
	@$(INFO) gofmt simplify
	@$(GOFMT) -l -s -w $(GO_SUBDIRS) $(GO_INTEGRATION_TESTS_SUBDIRS) || $(FAIL)
	@$(OK) gofmt simplify

go.validate: go.vet go.fmt

ifeq ($(GO111MODULE),on)

go.vendor.lite go.vendor.check:
	@$(INFO) verify dependencies have expected content
	@$(GOHOST) mod verify || $(FAIL)
	@$(OK) go modules dependencies verified

go.vendor.update:
	@$(INFO) update go modules
	@$(GOHOST) get -u ./... || $(FAIL)
	@$(OK) update go modules

go.vendor:
	@$(INFO) go mod vendor
	@$(GOHOST) mod vendor || $(FAIL)
	@$(OK) go mod vendor

else

go.vendor.lite: $(DEP)
#	dep ensure blindly updates the whole vendor tree causing everything to be rebuilt. This workaround
#	will only call dep ensure if the .lock file changes or if the vendor dir is non-existent.
	@if [ ! -d $(GO_VENDOR_DIR) ]; then \
		$(MAKE) vendor; \
	elif ! $(DEP) ensure -no-vendor -dry-run &> /dev/null; then \
		$(MAKE) vendor; \
	fi

go.vendor.check: $(DEP)
	@$(INFO) checking if vendor deps changed
	@$(DEP) check -skip-vendor || $(FAIL)
	@$(OK) vendor deps have not changed

go.vendor.update: $(DEP)
	@$(INFO) updating vendor deps
	@$(DEP) ensure -update -v || $(FAIL)
	@$(OK) updating vendor deps

go.vendor: $(DEP)
	@$(INFO) dep ensure
	@$(DEP) ensure || $(FAIL)
	@$(OK) dep ensure

endif

go.clean:
	@# `go modules` creates read-only folders under WORK_DIR
	@# make all folders within WORK_DIR writable, so they can be deleted
	@if [ -d $(WORK_DIR) ]; then chmod -R +w $(WORK_DIR); fi 

	@rm -fr $(GO_BIN_DIR) $(GO_TEST_DIR)

go.distclean:
	@rm -rf $(GO_VENDOR_DIR) $(GO_PKG_DIR)

go.generate:
	@$(INFO) go generate $(PLATFORM)
	@CGO_ENABLED=0 $(GOHOST) generate $(GO_GENERATE_FLAGS) $(GO_PACKAGES) $(GO_INTEGRATION_TEST_PACKAGES) || $(FAIL)
	@$(OK) go generate $(PLATFORM)
	@$(INFO) go mod tidy
	@$(GOHOST) mod tidy || $(FAIL)
	@$(OK) go mod tidy

.PHONY: go.init go.build go.install go.test.unit go.test.integration go.test.codecov go.lint go.vet go.fmt go.generate
.PHONY: go.validate go.vendor.lite go.vendor go.vendor.check go.vendor.update go.clean go.distclean

# ====================================================================================
# Common Targets

build.init: go.init
build.check: go.lint
build.code.platform: go.build
clean: go.clean
distclean: go.distclean
lint.init: go.init
lint.run: go.lint
test.init: go.init
test.run: go.test.unit
generate.init: go.init
generate.run: go.generate

# ====================================================================================
# Special Targets

fmt: go.imports
fmt.simplify: go.fmt.simplify
imports: go.imports
imports.fix: go.imports.fix
vendor: go.vendor
vendor.check: go.vendor.check
vendor.update: go.vendor.update
vet: go.vet

define GO_HELPTEXT
Go Targets:
    generate        Runs go code generation followed by goimports on generated files.
    fmt             Checks go source code for formatting issues.
    fmt.simplify    Format, simplify, update source files.
    imports         Checks go source code import lines for issues.
    imports.fix     Updates go source files to fix issues with import lines.
    vendor          Updates vendor packages.
    vendor.check    Fail the build if vendor packages have changed.
    vendor.update   Update vendor dependencies.
    vet             Checks go source code and reports suspicious constructs.
    test.unit.nocov Runs unit tests without coverage (faster for iterative development)
endef
export GO_HELPTEXT

go.help:
	@echo "$$GO_HELPTEXT"

help-special: go.help

.PHONY: fmt vendor vet go.help

# ====================================================================================
# Tools install targets

$(DEP):
	@$(INFO) installing dep-$(DEP_VERSION) $(SAFEHOSTPLATFORM)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-dep || $(FAIL)
	@curl -fsSL -o $(DEP) https://github.com/golang/dep/releases/download/$(DEP_VERSION)/dep-$(SAFEHOSTPLATFORM) || $(FAIL)
	@chmod +x $(DEP) || $(FAIL)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-dep
	@$(OK) installing dep-$(DEP_VERSION) $(SAFEHOSTPLATFORM)

$(GOLANGCILINT):
	@$(INFO) installing golangci-lint-v$(GOLANGCILINT_VERSION) $(SAFEHOSTPLATFORM)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-golangci-lint || $(FAIL)
	@curl -fsSL https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCILINT_VERSION)/golangci-lint-$(GOLANGCILINT_VERSION)-$(SAFEHOSTPLATFORM).tar.gz | tar -xz --strip-components=1 -C $(TOOLS_HOST_DIR)/tmp-golangci-lint || $(FAIL)
	@mv $(TOOLS_HOST_DIR)/tmp-golangci-lint/golangci-lint $(GOLANGCILINT) || $(FAIL)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-golangci-lint
	@$(OK) installing golangci-lint-v$(GOLANGCILINT_VERSION) $(SAFEHOSTPLATFORM)

$(GOFMT):
	@$(INFO) installing gofmt$(GOFMT_VERSION)
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-fmt || $(FAIL)
	@curl -sL https://dl.google.com/go/go$(GOFMT_VERSION).$(SAFEHOSTPLATFORM).tar.gz | tar -xz -C $(TOOLS_HOST_DIR)/tmp-fmt || $(FAIL)
	@mv $(TOOLS_HOST_DIR)/tmp-fmt/go/bin/gofmt $(GOFMT) || $(FAIL)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-fmt
	@$(OK) installing gofmt$(GOFMT_VERSION)

$(GOIMPORTS):
	@$(INFO) installing goimports
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-goimports || $(FAIL)
	@GO111MODULE=off GOPATH=$(TOOLS_HOST_DIR)/tmp-goimports GOBIN=$(TOOLS_HOST_DIR) $(GOHOST) get golang.org/x/tools/cmd/goimports || rm -fr $(TOOLS_HOST_DIR)/tmp-goimports || $(FAIL)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-goimports
	@$(OK) installing goimports

$(GOJUNIT):
	@$(INFO) installing go-junit-report
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-junit || $(FAIL)
	@GO111MODULE=off GOPATH=$(TOOLS_HOST_DIR)/tmp-junit GOBIN=$(TOOLS_HOST_DIR) $(GOHOST) get github.com/jstemmer/go-junit-report || rm -fr $(TOOLS_HOST_DIR)/tmp-junit || $(FAIL)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-junit
	@$(OK) installing go-junit-report

$(GOCOVER_COBERTURA):
	@$(INFO) installing gocover-cobertura
	@mkdir -p $(TOOLS_HOST_DIR)/tmp-gocover-cobertura || $(FAIL)
	@GO111MODULE=off GOPATH=$(TOOLS_HOST_DIR)/tmp-gocover-cobertura GOBIN=$(TOOLS_HOST_DIR) $(GOHOST) get github.com/t-yuki/gocover-cobertura || rm -fr $(TOOLS_HOST_DIR)/tmp-covcover-cobertura || $(FAIL)
	@rm -fr $(TOOLS_HOST_DIR)/tmp-gocover-cobertura
	@$(OK) installing gocover-cobertura
