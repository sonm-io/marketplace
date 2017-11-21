#!/bin/bash

GOROOT=${GOROOT:-`go env GOROOT`}
GO=${GOROOT}/bin/go
GOFMT=${GOROOT}/bin/gofmt

# This script MUST be run from the marketplace root.

# This script creates fake test files so that 'go test' generated correct reports for them
# and Cobertura in Jenkins could provide complete code coverage.

FAKE_TEST_FILE=fake_for_correct_coverage_test.go

${GO} list -f '{{.ImportPath}} {{.Name}}' ./... | # list all packages including path and package name (sometimes they differ)
    sed 's/^marketplace\///' |                     # trim `marketplace/` prefix
    grep -v 'vendor' |                         # exclude vendor
    grep -v '^mocks' |                         # exclude auto-generated mocks
    grep -v '^docs' |                          # exclude auto-generated docs
    grep -v '^tools' |                         # exclude tools - we're not going to additionally test them
    grep -v 'main$' |                          # exclude main package
    grep -v '^test/\?' |                       # exclude all folders under test/
    tail -n +2 |                               # exclude the root folder (e.g. marketplace/)
while read in; do
    import=`echo $in | cut -d " " -f 1`
    package=`echo $in | cut -d " " -f 2`
    ls ${import} | grep '_test.go$' > /dev/null
    if [ $? -eq 1 ]; then
        echo "Generated $import/$FAKE_TEST_FILE"
        echo "// DO NOT WRITE TESTS INTO THIS AUTO-GENERATED FILE, CREATE ANOTHER TEST FILE WITH A MEANINGFUL NAME" > ${import}/${FAKE_TEST_FILE}
        echo "package ${package}" >> ${import}/${FAKE_TEST_FILE}
    fi
done
