## Variables

UI_BUILD_DIR=$(BUILD_DIR)/ui
UI_DIR=$(WORKING_DIR)/ui

##@ Building - Frontend

.PHONY: frontend-build
frontend-build: $(BUILD_DIR)
	@echo "Building frontend..."
	@cd $(UI_DIR) && yarn install
	@cd $(UI_DIR) && yarn build
	@mv $(UI_DIR)/build $(BUILD_DIR)/ui

.PHONY: frontend-start
frontend-start:
	@echo "Starting development server..."
	@cd $(UI_DIR) && yarn start

.PHONY: frontend-start-api-mock
frontend-start-api-mock:
	@echo "Starting mock api..."
	cd $(UI_DIR) && npm install -g json-server
	cd $(UI_DIR)/api-mock && json-server --watch db.json --p 8080 --routes routes.json