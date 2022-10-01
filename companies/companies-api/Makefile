.DEFAULT_GOAL := help

.PHONY: help
help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## install dependencies
	@go get .
	@go mod download

lint-vet:
	@go vet ./...

lint: ## run lint
	@time golangci-lint run ./...
	@echo "Lint completed"

run: lint ## run app
	@go run main.go

build: lint ## build binary
	@go build -o app