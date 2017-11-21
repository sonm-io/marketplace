#!/bin/sh

PACKAGE_NAME=marketplace
PLATFORM=`uname`

find . -type f -name "*.go" -print0 | while IFS= read -r -d '' mock_file;
do

    if ! [[ $mock_file == *"/mocks/"* ]] || [[ $mock_file == *"/vendor/"* ]] || [[ $mock_file == *"/stub.go"* ]]
    then
       continue
    fi

    printf 'Cleaning mock: %s\n' "$mock_file";

    if [[ "$PLATFORM" == 'Darwin' ]]; then
        sed -i backup "s/${PACKAGE_NAME}\/vendor\///g" "${mock_file}"
        rm "${mock_file}backup"
    elif [[ "$PLATFORM" == 'Linux' ]]; then
        sed -i "s/${PACKAGE_NAME}\/vendor\///g" "${mock_file}"
    fi

done
