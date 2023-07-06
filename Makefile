.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t shota-imoto/otamesi2307:${DOCKER_TAG} \
	--target deploy ./
build-local: ## Build docker image to local development
	docker compose build --no-cache
up: ## Do docker compose up with hot reload
	docker compose up -d
down: ## Do docker compose down
	docker compose down
logs: ## Tail docker compose logs
	docker compose logs -f
ps: ## Check container status
	docker compose ps
test: ## Execute tests
	go test -shuffle=on ./...
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
