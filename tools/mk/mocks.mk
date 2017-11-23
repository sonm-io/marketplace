#ServiceInterfaces = $(shell egrep 'type\s+[a-zA-Z]+.*interface' ./*.go | cut -d" " -f 2 | xargs echo | sed 's/ /,/g' | sed -e "s/[[:space:]]//g")

.PHONY: clean-mocks
clean-mocks:
	bash tools/clean-mocks.sh

.PHONY: rm-mocks
rm-mocks:
	make clean-mocks | tail -n +4 | awk '{print $$3}' | xargs rm -rf
	rm -rf mocks.go

.PHONY: generate-mocks
generate-mocks: generate-mocks-intf-storage generate-mocks-usecase-marketplace-command clean-mocks
	@echo "Generating mocks"

.PHONY: generate-mocks-intf-storage
generate-mocks-intf-storage: tools
	@mkdir  -p ./interface/storage/mocks
	mockgen -package mocks \
            -source ./interface/storage/order.go Engine > ./interface/storage/mocks/order.go

.PHONY: generate-mocks-usecase-marketplace-command
generate-mocks-usecase-marketplace-command: tools
	@mkdir  -p ./usecase/marketplace/command/mocks
	mockgen -package mocks \
            -source ./usecase/marketplace/command/cancel_order_handler.go CancelOrderStorage > ./usecase/marketplace/command/mocks/cancel_order_storage.go
	mockgen -package mocks \
            -source ./usecase/marketplace/command/create_order_handler.go CreateOrderStorage > ./usecase/marketplace/command/mocks/create_order_storage.go
