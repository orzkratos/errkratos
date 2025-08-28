ERROR_PROTO_FILES=$(shell find errorspb -name *.proto)

.PHONY: install
# install errgenkratos plugin (since this project uses errgenkratos for code generation)
install:
	go install github.com/orzkratos/errgenkratos/cmd/protoc-gen-orzkratos-errors@latest
	@echo "success!"

.PHONY: generate
# generate errors code
generate:
	protoc --proto_path=./errorspb \
		   --proto_path=./proto3ps \
		   --go_out=paths=source_relative:./errorspb \
		   --orzkratos-errors_out=paths=source_relative:./errorspb \
		   $(ERROR_PROTO_FILES)

.PHONY: help
# show help
help:
	@echo "Available targets:"
	@echo "  install     - Install errgenkratos plugin"
	@echo "  generate    - Generate proto errors code"
	@echo "  help        - Show this help"
