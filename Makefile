DOCKER_BUILD := docker run --rm -u `id -u` -v ${PWD}:/sdk openapitools/openapi-generator-cli:v5.2.1 generate -i  sdk/api_files/postcast_api_openapi.json
GO_CLIENT := -g go -o /sdk/portcast \
			--git-repo-id=go-portcast-sdk --git-user-id=buyco \
			--additional-properties=packageName=portcast \
			--additional-properties=isGoSubmodule=false \
			--additional-properties=generateInterfaces=true

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

## generate: Clean and generate SDK from file.
generate:
	${MAKE} clean
	${MAKE} go-sdk

go-sdk:
	${DOCKER_BUILD} ${GO_CLIENT}

clean:
	rm -rf ./api

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo