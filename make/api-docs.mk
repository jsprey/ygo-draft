## Variables

API_FOLDER=$(WORKING_DIR)/backend/api
SWAG_BIN=$(UTILITY_BIN_PATH)/swag
SWAG_VERSION=v1.16.3

##@ Building - Backend

$(SWAG_BIN): $(UTILITY_BIN_PATH) ## Download swag.
	@echo "Download swag..."
	@$(call go-get-tool,$@,github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION))
	@echo "Downloaded swag successfully."

.PHONY: generate-api-docs
generate-api-docs: $(SWAG_BIN) ## Vet, test and build the backend.
	$(SWAG_BIN) init --output $(API_FOLDER)