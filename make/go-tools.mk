##@ Go-Utility Tools

GO_LINT=$(UTILITY_BIN_PATH)/golangci-lint
GO_LINT_VERSION=v1.45.2
$(GO_LINT): $(UTILITY_BIN_PATH) ## Download golangci-lint.
	@echo "Download golangci-lint..."
	@$(call go-get-tool,$@,github.com/golangci/golangci-lint/cmd/golangci-lint@$(GO_LINT_VERSION))

UTILITY_BIN_PATH=$(BUILD_DIR)/bin
$(UTILITY_BIN_PATH):
	@mkdir -p $@

# go-get-tool will 'go get' any package $2 and install it to $1.
define go-get-tool
	@[ -f $(1) ] || { \
		set -e ;\
		TMP_DIR=$$(mktemp -d) ;\
		cd $$TMP_DIR ;\
		go mod init tmp ;\
		echo "Downloading $(2)" ;\
		GOBIN=$(UTILITY_BIN_PATH) go install $(2) ;\
		rm -rf $$TMP_DIR ;\
	}
endef