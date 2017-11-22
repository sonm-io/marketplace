CYCLO_PACKAGES := $(foreach p,$(PACKAGES),cyclo-$(p))
CYCLO_AVG_PACKAGES := $(foreach p,$(PACKAGES),cyclo-avg-$(p))

.PHONY: tools-cyclo
tools-cyclo:
	${GO} get -u github.com/alecthomas/gocyclo

.PHONY: cyclo-avg
cyclo-avg: tools-cyclo
	@$(MAKE) $(CYCLO_AVG_PACKAGES) | grep -v NaN | sort -r -k 2 | head -n10

.PHONY: $(CYCLO_AVG_PACKAGES)
$(CYCLO_AVG_PACKAGES):
	$(eval $@_package := $(subst cyclo-avg-,,$@))
	@echo $($@_package) $$(gocyclo -avg $($@_package) | tail -n1 | awk '{ print $$2 }')

# Find duplicates in code
.PHONY: duplicates
duplicates: tools-dupl
	dupl -plumbing -threshold 50 | grep -v msgpack | grep -v test

LINT_EXCLUDE:= --exclude='.*/mocks/.*' \
    --exclude='.*_easyjson.go' \
    --exclude='.*_gen_test.go' \
    --exclude='.*_gen.go' \
    --exclude='.*bindata.go' \
    --exclude='.*/vendor/.*' \
    --exclude='.*_msgpack.go' \
    --exclude='.*_msgpack_test.go'

LINT_FAST := gometalinter $(LINT_EXCLUDE) \
    --deadline=3000s \
    --cyclo-over=50 \
    --min-const-length=4 \
    --min-occurrences=10 \
    --line-length=300 \
    --disable-all \
    --enable=vet \
    --enable=vetshadow \
    --enable=gosimple \
    --enable=staticcheck \
    --enable=ineffassign \
    --enable=gocyclo \
    --enable=lll \
    --enable=goconst \
    --linter='gohint:gohint -config=deploy/go_hint_config.json ./*.go:PATH:LINE:MESSAGE' --enable=gohint

LINT_SLOW := gometalinter $(LINT_EXCLUDE) \
    --deadline=3000s \
    --dupl-threshold=290 \
    --disable-all \
    --enable=unconvert \
    --enable=unused \
    --enable=varcheck \
    --enable=dupl \
    --enable=errcheck \
    --enable=deadcode \
    --enable=gas


.PHONY: lint-dep
lint-dep: #generate
	${GO} get -u github.com/elgris/hint/gohint
	${GO} get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: ci-lint
ci-lint: lint-dep
	$(LINT_FAST) --concurrency=4 --checkstyle --vendor ./... > checkstyle-result.xml || true
	if ! [ -z "$(RESULTS_DIR)" ] && [ -d "$(RESULTS_DIR)" ]; then cp checkstyle-result.xml "$(RESULTS_DIR)"/; fi

.PHONY: lint
lint: lint-dep
	$(LINT_FAST) --concurrency=4 --vendor ./... || true
	$(LINT_SLOW) --concurrency=2  --vendor ./... || true
