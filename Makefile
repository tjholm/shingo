include tools/tools.mk

build: init
	@echo building shingo server
	@go build -o bin/shingo cmd/main.go 

init: go-mod-download install-tools

go-mod-download:
	@echo installing go dependencies
	@go mod download

# Generate interfaces
generate-proto: install-tools
	@echo Generating Proto Sources
	@mkdir -p ./pkg/api/
	@$(PROTOC) --go_out=./pkg/api/ --validate_out="lang=go:./pkg/api" --go-grpc_out=./pkg/api -I ./proto ./proto/*/**/*.proto -I ./contracts
