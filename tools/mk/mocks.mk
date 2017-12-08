#ServiceInterfaces = $(shell egrep 'type\s+[a-zA-Z]+.*interface' ./*.go | cut -d" " -f 2 | xargs echo | sed 's/ /,/g' | sed -e "s/[[:space:]]//g")

.PHONY: clean-mocks
clean-mocks:
	@bash tools/clean-mocks.sh

.PHONY: rm-mocks
rm-mocks:
	make clean-mocks | tail -n +4 | awk '{print $$3}' | xargs rm -rf
	rm -rf mocks.go

.PHONY: generate-mocks
generate-mocks: generate-mocks-intf-storage generate-mocks-usecase-intf generate-mocks-usecase-marketplace-command generate-mocks-usecase-marketplace-query clean-mocks
	@echo "Generating mocks"

.PHONY: generate-mocks-intf-storage
generate-mocks-intf-storage: tools
	@mkdir  -p ./interface/storage/mocks
	@mockgen -package mocks \
            -source ./interface/storage/inmemory_engine.go Engine > ./interface/storage/mocks/engine.go

.PHONY: generate-mocks-usecase-marketplace-command
generate-mocks-usecase-marketplace-command: tools
	@mkdir  -p ./usecase/marketplace/command/mocks
	@mockgen -package mocks \
            -source ./usecase/marketplace/command/cancel_order_handler.go CancelOrderStorage > ./usecase/marketplace/command/mocks/cancel_order_storage.go
	@mockgen -package mocks \
            -source ./usecase/marketplace/command/create_bid_order_handler.go CreateBidOrderStorage > ./usecase/marketplace/command/mocks/create_bid_order_storage.go
	@mockgen -package mocks \
            -source ./usecase/marketplace/command/create_ask_order_handler.go CreateAskOrderStorage > ./usecase/marketplace/command/mocks/create_ask_order_storage.go

 .PHONY: generate-mocks-usecase-marketplace-query
generate-mocks-usecase-marketplace-query: tools
	@mkdir  -p ./usecase/marketplace/query/mocks
	@mockgen -package mocks \
            -source ./usecase/marketplace/query/get_order_handler.go OrderByIDStorage > ./usecase/marketplace/query/mocks/order_by_id_storage.go
	@mockgen -package mocks \
            -source ./usecase/marketplace/query/get_orders_handler.go OrderBySpecStorage > ./usecase/marketplace/query/mocks/order_by_spec_storage.go

.PHONY: generate-mocks-usecase-intf
generate-mocks-usecase-intf: tools
	@mkdir  -p ./usecase/intf/mocks
	@mockgen -package mocks \
			-destination ./usecase/intf/mocks/cqrs.go github.com/sonm-io/marketplace/usecase/intf Query,QueryHandler,Command,CommandHandler