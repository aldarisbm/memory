#!make
SHELL := /bin/bash

CURRENT_USER := $(shell whoami)

vet:
	go vet ./...

test:
	go test -v ./...

# should probably run a previous step checking for empty user instead of all in one
removefolder:
	if [ -z "$(CURRENT_USER)" ]; then \
		echo "CURRENT_USER is empty"; \
		exit 1; \
	fi

	rm -rf /Users/$(CURRENT_USER)/xyz.memorystore
