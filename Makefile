GO_REQ_VER := 1.9.2

# AVAILABLE ENV FROM CI
# - unit
# PATH            = $PATH:/usr/local/go/bin:$WORKSPACE/bin
# GOPATH          = $WORKSPACE
# SRC_API_DIR     = $WORKSPACE/src/local
# RESULTS_DIR     = $WORKSPACE/unit_results
# - integration
# PATH            = $PATH:/usr/local/go/bin:$WORKSPACE/bin
# GOPATH          = $WORKSPACE
# CATALOG_API_DIR = $WORKSPACE/src/marketplace
# RESULTS_DIR     = $WORKSPACE/integration_results
RESULTS_DIR := integration-results

# DO NOT TOUCH BELLOW
SHELL      = /bin/bash
BUILD_DIR := cli/cmd

# It assumes that GOPATH may be a list of paths
MAIN_GOPATH := $(lastword $(subst :, ,$(GOPATH)))
BIN         ?= $(MAIN_GOPATH)/bin/marketplace
BINDIR      ?= $(MAIN_GOPATH)/bin/

DATE        := $(shell date "+%Y-%m-%d %H:%M:%S")
GIT_LOG     := $(shell git log --decorate --oneline -n1| sed -e "s/'/ /g" -e "s/\"/ /g" -e "s/\#/\â„–/g" -e 's/`/ /g')
GIT_REV      := $(shell git rev-parse --short HEAD)
APP_VERSION := $(shell git branch|grep '*'| cut -f2 -d' ')

LDFLAGS = -X 'main.AppVersion=$(APP_VERSION)' \
          -X 'main.GitRev=$(GIT_REV)' \
          -X 'main.GoVersion=$(GO_REQ_VER)' \
          -X 'main.BuildDate=$(DATE)' \
          -X 'main.GitLog=$(GIT_LOG)'

GOFILES_NOVENDOR  = $(shell find . -type f -name '*.go' -not -path "*/vendor*")
PACKAGES         := $(shell find . -name '*.go' -not -path '*/vendor*' -not -path '*/test*' -exec dirname '{}' ';' | sort -u  | sed -e 's/^\.\///')

include ./tools/mk/tools.mk
include ./tools/mk/dep.mk
include ./tools/mk/lint.mk
include ./tools/mk/mocks.mk
include ./tools/mk/tests.mk

.PHONY: all
all: clean deps test fast-build
	@echo "Done"

.PHONY: clean
clean:
	rm -f $(BIN)
	rm -rf test-results
	rm -rf integration-results
#	find . -name "*_gen.go" | grep -v "handler2/docs" | (xargs rm || echo "Done")
#	find . -name "*_gen_test.go" | grep -v "handler2/docs" | xargs rm | (xargs rm || echo "Done")

.PHONY: deps
deps: get-dep
	@test -x $(DEP_BIN) || { echo "Dep $(DEP_REQ_VERSION) is required"; exit 1; }
	@mkdir -p ./vendor/
	$(DEP_BIN) ensure -v -vendor-only
	${GO} get github.com/golang/mock/mockgen

.PHONY: ci-deps
ci-deps: deps
	${GO} get github.com/axw/gocov/gocov
	${GO} get github.com/t-yuki/gocover-cobertura
	${GO} get -v github.com/tebeka/go2xunit
	${GO} get -v github.com/ttacon/chalk

.PHONY: build
build: clean deps fast-build

# fast build without check for deps
.PHONY: fast-build
fast-build:
	cd $(BUILD_DIR) && ${GO} build -i -ldflags "$(LDFLAGS)" -o $(BIN)
