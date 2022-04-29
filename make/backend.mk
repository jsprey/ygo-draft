## Variables

LDFLAGS=-ldflags "-extldflags -static -w -s -X main.Version=$(VERSION)"
GOOS=windows
GOARCH=amd64
ARTIFACT_BINARY_NAME=$(ARTIFACT_ID)-$(GOOS)-$(GOARCH)
GO_BUILD_FLAGS=-mod=vendor -a -tags netgo,osusergo $(LDFLAGS) -o $(BUILD_DIR)/$(ARTIFACT_BINARY_NAME)

##@ Building - Backend

.PHONY: build-backend
build-backend: $(BUILD_DIR) ## Builds the golang backend of YGODraft and creates a statically linked binary
	@echo "Compiling..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(GO_BUILD_FLAGS)