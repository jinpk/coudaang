#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

proto-gen:
	@echo "Generating Protobuf files"
	@sh ./scripts/protocgen.sh

.PHONY: proto-gen