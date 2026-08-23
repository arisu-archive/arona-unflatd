# Change these variables as necessary.
MAIN_PACKAGE_PATH := ./main.go
TMP_DIR := ./tmp
BINARY_NAME := arona-unflatd
GOLANGCI_LINT_VERSION := v2.12.2
GINKGO_VERSION := v2.28.3
GOVULNCHECK_VERSION := v1.7.0
GOLANGCI_LINT_FLAGS ?=

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

.PHONY: prepare
prepare:
	mkdir -p ${TMP_DIR}/bin
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}
	go install github.com/onsi/ginkgo/v2/ginkgo@${GINKGO_VERSION}
	go install golang.org/x/vuln/cmd/govulncheck@${GOVULNCHECK_VERSION}

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit: export CGO_ENABLED = 1
audit:
	go mod verify
	go vet ./...
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION} fmt --diff ./...
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION} run ${GOLANGCI_LINT_FLAGS} ./...
	go run golang.org/x/vuln/cmd/govulncheck@${GOVULNCHECK_VERSION} ./...

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run unit tests
.PHONY: test
test: export CGO_ENABLED = 1
test:
	ginkgo -r -cover -coverprofile=coverage.out

## build: build the application
.PHONY: build
build: export CGO_ENABLED = 1
build:
	go build -o=${TMP_DIR}/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the application
.PHONY: run
run: build
	${TMP_DIR}/bin/${BINARY_NAME}

# ==================================================================================== #
# OPERATIONS
# ==================================================================================== #

## push: push changes to the remote Git repository
.PHONY: push
push: tidy audit no-dirty
	git push

## production/deploy: deploy the application to production
.PHONY: production/build
production/build: export CGO_ENABLED = 1
production/build:
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o=${TMP_DIR}/bin/linux_amd64/${BINARY_NAME} ${MAIN_PACKAGE_PATH}
