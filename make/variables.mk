## Variables

WORKING_DIR=$(shell pwd)

# Contains all build artifacts
BUILD_DIR=$(WORKING_DIR)/build
$(BUILD_DIR):
	@echo "Creating build dir: $(BUILD_DIR)"
	@mkdir -p $@
