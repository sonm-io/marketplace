# Find packages with tests, excluding vendor
ROOT_PACKAGE                           := "github.com/sonm-io"
#ROOT_PACKAGE_ESC                       := $(echo "${ROOT_PACKAGE}" | sed -e 's/[]\/$*.^|[]/\\&/g')
FIND_UNIT_TEST_PACKAGES                := find . -name '*_test.go' -not -path '*/vendor*' -not -path '*/test*' -exec dirname '{}' ';' | sort -u
UNIT_TEST_PACKAGES                     := ${FIND_UNIT_TEST_PACKAGES} | sed -e 's/^\./github.com\/sonm-io\/marketplace/'
UNIT_TEST_PACKAGES_FOR_COVERAGE        := ${FIND_UNIT_TEST_PACKAGES} | sed -e 's/^\.\///'
UNIT_TEST_COVERAGE_TARGETS             := $(foreach p,$(shell ${UNIT_TEST_PACKAGES_FOR_COVERAGE}),unit-test-$(p))

FIND_INTEGRATION_TEST_PACKAGES         := find ./test -name '*_test.go' -not -path '*/vendor*' -exec dirname '{}' ';' | sort -u
INTEGRATION_TEST_PACKAGES              := ${FIND_INTEGRATION_TEST_PACKAGES} | sed -e 's/^\./github\.com\/sonm-io\/marketplace/'
INTEGRATION_TEST_PACKAGES_FOR_COVERAGE := ${FIND_INTEGRATION_TEST_PACKAGES} | sed -e 's/^\.\///'
INTEGRATION_TEST_COVERAGE_TARGETS      := $(foreach p,$(shell ${INTEGRATION_TEST_PACKAGES_FOR_COVERAGE}),integration-test-$(p))
.PHONY: create-fake-test-files
	create-fake-test-files:
	bash tools/generate-fake-tests.sh

test-results:
	mkdir integration-results
	mkdir -p test-results/cover
	mkdir -p test-results/output

.PHONY: test-setup
test-setup: tools
#Temporarily commented out setup of readdb
#	${GO} run test/setup/main.go
#	${GO} get github.com/jteeuwen/go-bindata/go-bindata && ${GO} install github.com/jteeuwen/go-bindata/go-bindata

.PHONY: test
test: unit-test integration-test

.PHONY: ci-test-config
ci-test-config:
#	if [ ! -f etc/test.ini ]; then cp etc/{_,}test.ini; fi
#	if [ -f _test_credentials.yaml ]; then cp {_,}test_credentials.yaml; fi

.PHONY: ci-test-all
ci-test-all: unit-test ci-test-config integration-test
ci-test-all:
	if ! [ -z "$(RESULTS_DIR)" ] && [ -d "$(RESULTS_DIR)" ]; then \
	    cat test-results/output/* | go2xunit -gocheck -fail -fail-on-race -output "$(RESULTS_DIR)"/report.xml; \
	    echo "mode: set" | cat - report.out | gocover-cobertura > "$(RESULTS_DIR)"/coverage.xml; \
	fi

.PHONY: test-check
test-check:
	${UNIT_TEST_PACKAGES} | xargs ${GO} test -run=^bad
	${INTEGRATION_TEST_PACKAGES} | xargs ${GO} test -run=^bad

.PHONY: fast-unit-test
fast-unit-test:
	${UNIT_TEST_PACKAGES} | xargs ${GO} test -timeout 10s

.PHONY: fast-unit-test-fmt
fast-unit-test-fmt:
	time ${UNIT_TEST_PACKAGES} | xargs ${GO} test -timeout 10s | column -t

.PHONY: race-unit-test
race-unit-test:
	${UNIT_TEST_PACKAGES} | xargs ${GO} test -timeout 10s -race

.PHONY: unit-test
unit-test: $(UNIT_TEST_COVERAGE_TARGETS)

.PHONY: $(UNIT_TEST_COVERAGE_TARGETS)
$(UNIT_TEST_COVERAGE_TARGETS): generate tools test-results create-fake-test-files race-unit-test
	$(eval $@_package := $(subst unit-test-,,$@))
	$(eval $@_coverprofile := $(subst /,_,$($@_package)))

	# tests are tee'ed for further parsing in Jenkins.
	set -o pipefail; ${GO} test -test.v -timeout 10s -coverprofile test-results/cover/$($@_coverprofile).out ${ROOT_PACKAGE}/marketplace/$($@_package) 2>&1 \
	  | tee test-results/output/unit_$($@_coverprofile)

	# remove lines from black list
	#sed -i '/_easyjson.go/d' test-results/cover/$($@_coverprofile).out
	#sed -i '/mocks/d'        test-results/cover/$($@_coverprofile).out
	#sed -i '/_gen.go/d'      test-results/cover/$($@_coverprofile).out
	#sed -i '/_msgpack.go/d'  test-results/cover/$($@_coverprofile).out

	# collect all cover reports without mode (default it is "mode: set")
	sed '1d' test-results/cover/$($@_coverprofile).out >> report.out || true

# Integration tests are those started with 'test/'.
# Note that if you omit the slash (\), bash variables won't be interpolated.

.PHONY: fast-integration-test
fast-integration-test:
	${INTEGRATION_TEST_PACKAGES} | xargs -I % ${GO} test -timeout 5m % -check.vv

.PHONY: race-integration-test
race-integration-test:
	${INTEGRATION_TEST_PACKAGES} | xargs -I % ${GO} test -timeout 5m -race % -check.vv

.PHONY: integration-test
integration-test: $(INTEGRATION_TEST_COVERAGE_TARGETS)

.PHONY: $(INTEGRATION_TEST_COVERAGE_TARGETS)
$(INTEGRATION_TEST_COVERAGE_TARGETS): generate-msgp test-setup test-results race-integration-test
	$(eval $@_package := $(subst integration-test-,,$@))
	$(eval $@_coverprofile := $(subst /,_,$($@_package)))

# List test imports of the package.
# grep only marketplace packages.
# Exclude test packages.
# Concatenate them with a comma (,) into a single line.
# Remove the trailing comma.
# Example of the result: "marketplace/common,marketplace/service"
# Tests are tee'ed for further parsing in Jenkins.
	@coverpkg=`${GO} list -f '{{.TestImports}}' ./$($@_package) \
	| grep -o -E 'marketplace/[^ ]+' | grep -v '^marketplace/test' | grep -v '/vendor/' \
	| tr '\n' ',' | sed 's/,$$//'`; \
	${GO} test -i marketplace/$($@_package); \
if [ -z $${coverpkg} ]; then \
	set -o pipefail; ${GO} test -test.v -timeout 2m marketplace/$($@_package) -check.vv 2>&1 \
	  | tee test-results/output/integration_$($@_coverprofile); \
else \
	set -o pipefail; ${GO} test -test.v -timeout 2m -coverprofile test-results/cover/$($@_coverprofile).out \
	  -coverpkg=$${coverpkg} marketplace/$($@_package) -check.vv 2>&1 \
	  | tee test-results/output/integration_$($@_coverprofile); \
	sed -i '/_easyjson.go/d' test-results/cover/$($@_coverprofile).out; \
	sed -i '/mocks/d'        test-results/cover/$($@_coverprofile).out; \
	sed -i '/_gen.go/d'      test-results/cover/$($@_coverprofile).out; \
	sed -i '/_msgpack.go/d'  test-results/cover/$($@_coverprofile).out; \
	sed '1d' test-results/cover/$($@_coverprofile).out >> report.out || true; \
fi;

# Errors like "warning: no packages being tested depend on marketplace/dao/datastruct" may occur. It'a alright
# and results from the way we run integration tests. No "real code" is located under folders like "marketplace/test/service"
# and golang consider this as a warning, see https://golang.org/src/cmd/go/test.go for more details.
