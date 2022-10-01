DOCKER_TAG=companies:local

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

docker-build: ## build docker image
	@docker build . -t $(DOCKER_TAG)

docker-sh: docker-build ## sh on docker image
	@docker run -it --rm --entrypoint /bin/sh $(DOCKER_TAG)

docker-run: docker-build ## run docker image
	@docker run -it --rm -p 80:80 $(DOCKER_TAG)

docker-build-builder:
	@docker build . -t $(DOCKER_TAG)-builder --target builder

docker-sh-builder: docker-build-builder
	@docker run -it --rm --entrypoint /bin/sh $(DOCKER_TAG)-builder

dc-up-dependencies: ## start docker-compose dependencies
	@docker-compose up -d companies-db

dc-stop: ## stop docker compose containers
	@docker-compose stop -t 0