## Variables

##@ Docker Compose - Dependencies

.PHONY: dc-down
dc-down: ## Shuts down the docker compose dependencies.
	docker compose down

.PHONY: dc-delete
dc-delete: ## Shuts down the docker compose dependencies and removes all volumes.
	docker compose down -v

.PHONY: dc-up
dc-up: ## Starts the docker compose dependencies.
	docker compose up -d