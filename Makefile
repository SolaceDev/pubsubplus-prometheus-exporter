#make
PROJECT_NAME := "pubsubplus-prometheus-exporter"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)


.PHONY:
dep: ## Get the dependencies
	@go mod vendor

.PHONY:
lint: ## Lint Golang files
	@./bin/golangci-lint run ./...  -D errcheck  #disable errcheck for now

.PHONY:
vet: ## Run go vet
	@go vet ${PKG_LIST}

.PHONY:
test: ## Run unit tests
	@go test -short ${PKG_LIST}

.PHONY:
test-coverage: ## Run tests with coverage
	@go test -short -coverprofile reports/cover.out ${PKG_LIST}
	@go tool cover -html reports/cover.out -o reports/cover.html

.PHONY:
build: dep ## Build the binary file
	@go build -a -ldflags '-s -w -extldflags "-static"' -o solace_prometheus_exporter

.PHONY:
clean: ## Remove previous build
	@rm -f cover.html cover.out solace_prometheus_exporter

# Generate third party license
.PHONY:
generate-license: dep ## Generate third party license
	@go install github.com/google/go-licenses@latest
	@go-licenses report . --template ci/whitesource/license-template.tpl > THIRD-PARTY-LICENSES.md

.PHONY:
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
