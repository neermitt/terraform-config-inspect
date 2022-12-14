TEST?=$$(go list ./... | grep -v 'vendor')
SHELL := /bin/bash
#GOOS=darwin
GOOS=linux
GOARCH=amd64
VERSION=test

# List of targets the `readme` target should call before generating the readme
export README_DEPS ?= docs/targets.md

-include $(shell curl -sSL -o .build-harness "https://cloudposse.tools/build-harness"; echo .build-harness)

get:
	go get

build: get
	env GOOS=${GOOS} GOARCH=${GOARCH} go build -o build/terraform-config-inspect -v -ldflags "-X 'github.com/neermitt/terraform-config-inspect/cmd.Version=${VERSION}'"

version: build
	chmod +x ./build/terraform-config-inspect
	./build/terraform-config-inspect --version

deps:
	go mod download

# Run acceptance tests
testacc: get
	go test $(TEST) -v $(TESTARGS) -timeout 2m

.PHONY: get build deps version testacc
