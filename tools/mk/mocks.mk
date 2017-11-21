ServiceInterfaces = $(shell egrep 'type\s+[a-zA-Z]+.*interface' ./*.go | cut -d" " -f 2 | xargs echo | sed 's/ /,/g' | sed -e "s/[[:space:]]//g")

.PHONY: clean-mocks
clean-mocks:
	bash tools/clean-mocks.sh

.PHONY: rm-mocks
rm-mocks:
	make clean-mocks | tail -n +4 | awk '{print $$3}' | xargs rm -rf
	rm -rf mocks.go

.PHONY: generate-mocks
generate-mocks: generate-mocks-service clean-mocks
	@echo "Generating mocks"

.PHONY: generate-mocks-service
generate-mocks-service: tools
	mockgen -package marketplace \
            -source ./server.go marketplace > mocks.go
