# Set these to the desired values
ARTIFACT_ID=ygodraft
VERSION=0.0.1

include make/variables.mk
include make/go-tools.mk
include make/backend.mk
include make/frontend.mk

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## Clean up all artifact data
	@rm -rf $(BUILD_DIR)

##@ Building

.PHONY: build
build: build-backend build-frontend ## Builds the backend and frontend of YGODraft.