SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache
 
 ## varable for `go`
GOCMD=go 
 ## variable  `build`
BUILD_DIR=build
# variable as build/bin
BINARY_DIR=$(BUILD_DIR)/bin

 # to start the application
run:
	@echo "Welcome To My Ecommerce Project"
	$(GOCMD) run ./cmd/api

 ## Generate swagger docs
swag: 
	swag init -g pkg/api/server.go -o ./cmd/api/docs
 
## Display this help screen
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
