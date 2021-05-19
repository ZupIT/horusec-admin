APP_NAME=horusec-admin
DOCKER_REPO=docker.io/horuszup
VERSION=1.0.0
ENVIRONMENT=production
IMG ?= $(DOCKER_REPO)/$(APP_NAME):$(VERSION)
GO ?= go
GOFMT ?= gofmt
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOCILINT ?= ./bin/golangci-lint
KUSTOMIZE = $(shell pwd)/bin/kustomize
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

kustomize: ## Download kustomize locally if necessary
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v3@v3.8.7)

build: ## Build the container
	docker build -t $(IMG) . -f ./deployments/Dockerfile

run: ## Run container on port 8007
	docker run -i -t --rm -p=8007:3000 --name="$(APP_NAME)" $(IMG)

install:
	cd deployments/horusec-terraform && make up

stop: ## Stop and remove a running container
	docker stop $(APP_NAME)

publish: ## Publish the container to Docker Hub
	docker push $(IMG)

deploy: kustomize ## Deploy horusec-admin in the configured Kubernetes cluster in ~/.kube/config
	cd $(PROJECT_DIR)/deployments/k8s/overlays/$(ENVIRONMENT); $(KUSTOMIZE) edit set image $(IMG)
	$(KUSTOMIZE) build deployments/k8s/overlays/$(ENVIRONMENT) | kubectl apply -f -

undeploy: ## UnDeploy horusec-admin from the configured Kubernetes cluster in ~/.kube/config
	$(KUSTOMIZE) build deployments/k8s/overlays/$(ENVIRONMENT) | kubectl delete -f -

fmt: ## Format all Go files
	$(GOFMT) -w $(GOFMT_FILES)

coverage: ## Run converage with threshold
	chmod +x deployments/scripts/coverage.sh
	deployments/scripts/coverage.sh 99 "./..."

lint: ## Run lint checks
    ifeq ($(wildcard $(GOCILINT)), $(GOCILINT))
		$(GOCILINT) run --timeout=5m -c .golangci.yml ./...
    else
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.39.0
		$(GOCILINT) run --timeout=5m -c .golangci.yml ./...
    endif

# go-get-tool will 'go get' any package $2 and install it to $1.
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
