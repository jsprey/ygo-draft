## Variables

LDFLAGS=-ldflags "-extldflags -static -w -s -X main.Version=$(VERSION)"
GOOS=linux
GOARCH=amd64
ARTIFACT_BINARY_NAME=$(ARTIFACT_ID)-$(GOOS)-$(GOARCH)
GO_BUILD_FLAGS=-a -tags netgo,osusergo $(LDFLAGS) -o $(BUILD_DIR)/$(ARTIFACT_BINARY_NAME)

ARTIFACT_TEST_JUNIT_REPORT=$(BUILD_DIR)/unit-report.xml

##@ Building - Backend

.PHONY: b-check
b-check: b-vet b-unit-test b-build ## Vet, test and build the backend.
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GO_BUILD_FLAGS)

.PHONY: b-run
b-run: generate-api-docs ## Run the go backend.
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go run .

.PHONY: b-run-air
b-run-air: $(GO_AIR) generate-api-docs ## Run the go backend.
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_AIR) .

.PHONY: b-build
b-build: $(BUILD_DIR) generate-api-docs ## Builds the golang backend of YGODraft and creates a statically linked binary.
	@echo "Running go build..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GO_BUILD_FLAGS)

.PHONY: b-unit-test
b-unit-test: $(BUILD_DIR) ## Runs all unit test for the backend.
	@echo "Running go test..."
	go test ./... -v

.PHONY: b-vet
b-vet: $(BUILD_DIR) ## Runs go vet on the backend code.
	@echo "Running go vet..."
	@go vet .