## Variables
##@ Building - Frontend

.PHONY: build-frontend
build-frontend:
	@echo "Building frontend..."

.PHONY: clean-ui-build
clean-ui-build:
	rm -rf ui/build
	rm -rf ui/node_modules
	rm -rf target/ui

target/ui/build: ui/build
	mkdir -p target/ui
	cp -r ui/build target/ui/build

ui/build: ui/src ui/public ui/node_modules
	cd ui && yarn build

ui/node_modules: ui/yarn.lock ui/package.json
	cd ui && yarn install
