GO_SYS_VER:=$(shell go version | cut -d" " -f3 | sed 's/go//')
GO_SYS_BIN:=$(shell which go)


# try to find Go if not defined into upper-level Makefile
# or via the env variables
ifeq ($(GO), )
	GO:=$(shell which go)
endif

GOFMT:=$(shell which gofmt)

$(info checking version of $(GO)…)
ifneq ($(GO_SYS_VER), $(GO_REQ_VER))
$(info -- found GO ${GO_SYS_VER} at ${GO_SYS_BIN}, but required version is ${GO_REQ_VER}. Use custom build...)
#	_=$(shell ./tools/install_go.sh $(GO_REQ_VER) $(INSTALL_PATH))
#	GO=$(INSTALL_PATH)/bin/go
#	export GOROOT:=${INSTALL_PATH}
#	export PATH:=${INSTALL_PATH}/bin:${PATH}
else
$(info -- found GO $(GO_SYS_VER) at ${GO_SYS_BIN})
endif

.PHONY: generate
#generate: generate-bindata generate-mocks generate-msgp generate-easyjson
generate: generate-mocks

.PHONY: fmt
fmt:
	${GOFMT} -s -w $(GOFILES_NOVENDOR)
