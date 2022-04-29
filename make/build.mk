## Includes

include make/build_backend.mk
include make/build_frontend.mk

##@ Building

.PHONY: build
build: build-backend build-frontend ## Builds the backend and frontend of YGODraft.