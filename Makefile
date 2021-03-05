APP_NAME=horusec-admin
DOCKER_REPO=docker.io/horuszup
GO ?= go
GOFMT ?= gofmt
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOCILINT ?= ./bin/golangci-lint

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build the container
	docker build -t $(DOCKER_REPO)/$(APP_NAME):v2 .

run: ## Run container on port 3000
	docker run -i -t --rm -p=3000:3000 --name="$(APP_NAME)" $(DOCKER_REPO)/$(APP_NAME):v2

stop: ## Stop and remove a running container
	docker stop $(APP_NAME)

publish: ## Publish the `v2` container to Docker Hub
	@echo 'publish v2 to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(APP_NAME):v2


fmt: # Format all Go files
	$(GOFMT) -w $(GOFMT_FILES)

# Run converage with threshold
coverage:
	chmod +x deployments/scripts/coverage.sh
	deployments/scripts/coverage.sh 99 "./..."

lint: ## Run lint checks
    ifeq ($(wildcard $(GOCILINT)), $(GOCILINT))
		$(GOCILINT) run -v --timeout=5m -c .golangci.yml ./...
    else
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.25.0
		$(GOCILINT) run -v --timeout=5m -c .golangci.yml ./...
    endif
