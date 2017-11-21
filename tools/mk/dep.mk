DEP_REQ_VERSION:=v0.3.2

DEP_CUSTOM_PATH:=$(GOPATH)/src/github.com/golang/dep
DEP_CUSTOM_BIN:=$(DEP_CUSTOM_PATH)/dep-$(DEP_REQ_VERSION)

DEP_SYS_INSTALLED:=$(shell command -v dep 2> /dev/null)
DEP_CUSTOM_INSTALLED:=$(shell command -v $(DEP_CUSTOM_BIN) 2> /dev/null)

DEP_BIN:=false
DEP_SYS_VER:=false
DEP_CUSTOM_VER:=false

ifdef DEP_SYS_INSTALLED
    DEP_SYS_VER:=$(shell dep version | sed -n 2p | cut -d":" -f2 | xargs)
endif

ifdef DEP_CUSTOM_INSTALLED
	DEP_CUSTOM_VER:=$(shell $(DEP_CUSTOM_BIN) version | sed -n 2p | cut -d":" -f2 | xargs)
endif

ifeq ($(DEP_REQ_VERSION),$(DEP_SYS_VER))
	DEP_BIN:=$(shell which dep)
else ifeq ($(DEP_REQ_VERSION),$(DEP_CUSTOM_VER))
	DEP_BIN:=$(DEP_CUSTOM_BIN)
endif

.PHONY: get-dep
get-dep:
ifeq ($(DEP_BIN),false)
	$(info #Installing dep version $(DEP_REQ_VERSION)...)

ifeq ($(wildcard $(DEP_CUSTOM_PATH)),)
	mkdir -p $(DEP_CUSTOM_PATH) && cd $(DEP_CUSTOM_PATH); \
	git clone https://github.com/golang/dep.git .
endif

ifndef DEP_CUSTOM_INSTALLED
	cd $(DEP_CUSTOM_PATH) && git fetch --tags && git checkout $(DEP_REQ_VERSION) && git reset --hard HEAD; \
	$(GO) build -a -installsuffix cgo -ldflags "-s -w -X main.version=$(DEP_REQ_VERSION)" -o $(DEP_CUSTOM_BIN) ./cmd/dep
endif

	$(eval DEP_BIN := $(DEP_CUSTOM_BIN))
	$(eval DEP_CUSTOM_VER := $(shell $(DEP_CUSTOM_BIN) version | sed -n 2p | cut -d":" -f2 | xargs))

endif
	$(DEP_BIN) version
