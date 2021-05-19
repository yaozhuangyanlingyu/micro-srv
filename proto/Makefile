.PHONY: proto
PARENT_DIR = $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../)

proto_path:
	@echo $(PARENT_DIR)

# cart-srv
cart_go_proto:
	protoc --proto_path=$(PARENT_DIR) --micro_out=$(PARENT_DIR) --go_out=$(PARENT_DIR) $(PARENT_DIR)/proto/cart/*.proto

cart_php_proto:
	protoc --proto_path=$(PARENT_DIR) --php_out=$(PARENT_DIR)/proto/php --grpc_out=$(PARENT_DIR)/proto/php --plugin=protoc -gen-grpc=$(GRPC_PATH) $(PARENT_DIR)/proto/cart/*.proto


# account-srv
account_go_proto:
	protoc --proto_path=$(PARENT_DIR) --micro_out=$(PARENT_DIR) --go_out=$(PARENT_DIR) $(PARENT_DIR)/proto/account/*.proto
