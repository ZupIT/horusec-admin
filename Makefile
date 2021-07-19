APP_NAME=horusec-admin
ENVIRONMENT=production
GO_IMPORTS ?= goimports
GO_IMPORTS_LOCAL ?= github.com/ZupIT/horusec-admin
ADMIN_VERSION ?= $(shell semver get release)
REGISTRY_IMAGE ?= horuszup/$(APP_NAME):$(ADMIN_VERSION)
GO ?= go
GOFMT ?= gofmt
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOCILINT ?= ./bin/golangci-lint
CONTROLLER_GEN ?= $(shell pwd)/bin/controller-gen
KUSTOMIZE = $(shell pwd)/bin/kustomize
HORUSEC ?= horusec
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
ADDLICENSE ?= addlicense

.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

install-semver: # Install semver binary
	curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec-devkit/main/scripts/install-semver.sh | bash

kustomize: ## Download kustomize locally if necessary
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v3@v3.8.7)

controller-gen:
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1)

generate: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

docker-up-alpha: ## Update alpha in docker image
	chmod +x ./deployments/scripts/update-image.sh
	./deployments/scripts/update-image.sh alpha false

docker-up-rc: ## Update rc in docker image
	chmod +x ./deployments/scripts/update-image.sh
	./deployments/scripts/update-image.sh rc false

docker-up-release: ## Update release in docker image
	chmod +x ./deployments/scripts/update-image.sh
	./deployments/scripts/update-image.sh release false

docker-up-release-latest: ## Update release and latest in docker image
	chmod +x ./deployments/scripts/update-image.sh
	./deployments/scripts/update-image.sh release true

docker-up-minor-latest: ## Update minor and latest in docker image
	chmod +x ./deployments/scripts/update-image.sh
	./deployments/scripts/update-image.sh minor true

run: install-semver ## Run container on port 8007
	docker run -i -t --rm -p=8007:3000 --name="$(APP_NAME)" $(REGISTRY_IMAGE)

run-dev:
	go run ./cmd/app/

stop: ## Stop and remove a running container
	docker stop $(APP_NAME)

generate-service-yaml: kustomize install-semver
	mkdir -p $(shell pwd)/tmp
	cd $(PROJECT_DIR)/deployments/k8s/overlays/$(ENVIRONMENT); $(KUSTOMIZE) edit set image $(REGISTRY_IMAGE)
	$(KUSTOMIZE) build deployments/k8s/overlays/$(ENVIRONMENT) > $(shell pwd)/tmp/horusec-admin.yaml

deploy: kustomize install-semver ## Deploy horusec-admin in the configured Kubernetes cluster in ~/.kube/config
	cd $(PROJECT_DIR)/deployments/k8s/overlays/$(ENVIRONMENT); $(KUSTOMIZE) edit set image $(REGISTRY_IMAGE)
	$(KUSTOMIZE) build deployments/k8s/overlays/$(ENVIRONMENT) | kubectl apply -f -

undeploy: ## UnDeploy horusec-admin from the configured Kubernetes cluster in ~/.kube/config
	$(KUSTOMIZE) build deployments/k8s/overlays/$(ENVIRONMENT) | kubectl delete -f -

fmt: ## Format all Go files
	$(GOFMT) -w $(GOFMT_FILES)

coverage: ## Check coverage in application
	curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec-devkit/main/scripts/coverage.sh | bash -s 0 .

build: ## Check coverage in application
	$(GO) build -o "./tmp/bin/admin" ./cmd/app

security: # Run security pipeline
    ifeq (, $(shell which $(HORUSEC)))
		curl -fsSL https://raw.githubusercontent.com/ZupIT/horusec/master/deployments/scripts/install.sh | bash -s latest
		$(HORUSEC) start -p="./" -e="true"
    else
		$(HORUSEC) start -p="./" -e="true"
    endif

lint: ## Run lint checks
    ifeq ($(wildcard $(GOCILINT)), $(GOCILINT))
		$(GOCILINT) run --timeout=5m -c .golangci.yml ./...
    else
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.39.0
		$(GOCILINT) run --timeout=5m -c .golangci.yml ./...
    endif

fix-imports: # Setup all imports to default mode
    ifeq (, $(shell which $(GO_IMPORTS)))
		$(GO) get -u golang.org/x/tools/cmd/goimports
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GOFMT_FILES)
    else
		$(GO_IMPORTS) -local $(GO_IMPORTS_LOCAL) -w $(GOFMT_FILES)
    endif

license:
	$(GO) get -u github.com/google/addlicense
	@$(ADDLICENSE) -check -f ./copyright.txt $(shell find -regex '.*\.\(go\|js\|ts\|yml\|yaml\|sh\|dockerfile\)')

license-fix:
	$(GO) get -u github.com/google/addlicense
	@$(ADDLICENSE) -f ./copyright.txt $(shell find -regex '.*\.\(go\|js\|ts\|yml\|yaml\|sh\|dockerfile\)')

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
