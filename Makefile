.DEFAULT_GOAL := build
.PHONY: build
BUILDFLAGS := -mod=vendor -ldflags="-s -w" -gcflags=-trimpath=$(CURDIR)

build:
	go build $(BUILDFLAGS) ./cmd/razweb
