ERROR_PROTO_FILES=$(shell find errorspb -name *.proto)

.PHONY: install
install:
	go install github.com/orzkratos/errgenkratos/cmd/protoc-gen-orzkratos-errors@latest
	@echo "success!"

.PHONY: generate
generate:
	protoc --proto_path=./errorspb \
		   --proto_path=./proto3ps \
		   --go_out=paths=source_relative:./errorspb \
		   --orzkratos-errors_out=paths=source_relative:./errorspb \
		   $(ERROR_PROTO_FILES)
	@echo "complete!"

.PHONY: clean
clean:
	rm -f errorspb/*.pb.go errorspb/*_errors.pb.go
	@echo "done!"

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  install  - Install protoc-gen-orzkratos-errors plugin"
	@echo "  generate - Generate Go code from proto files"
	@echo "  clean    - Remove generated .pb.go files"
	@echo "  help     - Show this help message"
