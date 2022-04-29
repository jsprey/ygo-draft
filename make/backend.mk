## Variables

LDFLAGS=-ldflags "-extldflags -static -w -s -X main.Version=$(VERSION)"
GOOS=linux
GOARCH=amd64
ARTIFACT_BINARY_NAME=$(ARTIFACT_ID)-$(GOOS)-$(GOARCH)
GO_BUILD_FLAGS=-a -tags netgo,osusergo $(LDFLAGS) -o $(BUILD_DIR)/$(ARTIFACT_BINARY_NAME)

ARTIFACT_TEST_JUNIT_REPORT=$(BUILD_DIR)/unit-report.xml

##@ Building - Backend

.PHONY: backend-check
backend-check: backend-vet backend-unit-test backend-build ## Vet, test and build the backend.
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GO_BUILD_FLAGS)

.PHONY: backend-build
backend-build: $(BUILD_DIR) ## Builds the golang backend of YGODraft and creates a statically linked binary.
	@echo "Running go build..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GO_BUILD_FLAGS)

.PHONY: backend-unit-test
backend-unit-test: $(BUILD_DIR) ## Runs all unit test for the backend.
	@echo "Running go test..."
	go test ./... -v

.PHONY: backend-vet
backend-vet: $(BUILD_DIR) ## Runs go vet on the backend code.
	@echo "Running go vet..."
	@go vet .