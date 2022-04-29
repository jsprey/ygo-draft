## Variables

WORK_DIR=$(pwd)

# Contains all build artifacts
BUILD_DIR=$(WORK_DIR)/build
$(BUILD_DIR):
	@echo "Creating build dir: $(BUILD_DIR)"
	@mkdir build
