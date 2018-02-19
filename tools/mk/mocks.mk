#ServiceInterfaces = $(shell egrep 'type\s+[a-zA-Z]+.*interface' ./*.go | cut -d" " -f 2 | xargs echo | sed 's/ /,/g' | sed -e "s/[[:space:]]//g")

.PHONY: clean-mocks
clean-mocks:
	@bash tools/clean-mocks.sh

.PHONY: rm-mocks
rm-mocks:
	make clean-mocks | tail -n +4 | awk '{print $$3}' | xargs rm -rf
	rm -rf mocks.go

.PHONY: generate-mocks
generate-mocks: generate-mocks-handler generate-mocks-service clean-mocks
	@echo "Generating mocks"

.PHONY: generate-mocks-service
generate-mocks-service: tools
	@mkdir  -p ./service/mocks
	@mockgen -package mocks \
            -source ./service/service.go Storage > ./service/mocks/storage.go

.PHONY: generate-mocks-handler
generate-mocks-handler: tools
	@mkdir  -p ./handler/srv/mocks
	@mockgen -package mocks \
            -source ./handler/srv/marketplace.go MarketService > ./handler/srv/mocks/market_service.go
