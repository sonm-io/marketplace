#ServiceInterfaces = $(shell egrep 'type\s+[a-zA-Z]+.*interface' ./*.go | cut -d" " -f 2 | xargs echo | sed 's/ /,/g' | sed -e "s/[[:space:]]//g")

.PHONY: clean-mocks
clean-mocks:
	@bash tools/clean-mocks.sh

.PHONY: rm-mocks
rm-mocks:
	make clean-mocks | tail -n +4 | awk '{print $$3}' | xargs rm -rf
	rm -rf mocks.go

.PHONY: generate-mocks
generate-mocks: generate-mocks-intf-reporting generate-mocks-usecase-intf generate-mocks-usecase-marketplace-command clean-mocks
	@echo "Generating mocks"

generate-mocks-intf-reporting: tools
	@mkdir  -p ./interface/reporting/sqllite/mocks
	@mockgen -package mocks \
            -source ./interface/reporting/sqllite/order_by_id_handler.go OrderRowFetcher > ./interface/reporting/sqllite/mocks/order_row_fetcher.go
	@mockgen -package mocks \
            -source ./interface/reporting/sqllite/match_orders_handler.go OrderRowsFetcher > ./interface/reporting/sqllite/mocks/order_rows_fetcher.go

.PHONY: generate-mocks-usecase-marketplace-command
generate-mocks-usecase-marketplace-command: tools
	@mkdir  -p ./usecase/marketplace/command/mocks
	@mockgen -package mocks \
            -source ./usecase/marketplace/command/cancel_order_handler.go OrderCanceler > ./usecase/marketplace/command/mocks/order_canceler.go
	@mockgen -package mocks \
            -source ./usecase/marketplace/command/create_bid_order_handler.go CreateBidOrderStorage > ./usecase/marketplace/command/mocks/create_bid_order_storage.go
	@mockgen -package mocks \
            -source ./usecase/marketplace/command/create_ask_order_handler.go CreateAskOrderStorage > ./usecase/marketplace/command/mocks/create_ask_order_storage.go
	@mockgen -package mocks \
            -source ./usecase/marketplace/command/touch_orders_handler.go OrderToucher > ./usecase/marketplace/command/mocks/order_toucher.go

.PHONY: generate-mocks-usecase-intf
generate-mocks-usecase-intf: tools
	@mkdir  -p ./usecase/intf/mocks
	@mockgen -package mocks \
			-destination ./usecase/intf/mocks/cqrs.go github.com/sonm-io/marketplace/usecase/intf Query,QueryHandler,Command,CommandHandler