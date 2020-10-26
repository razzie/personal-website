.DEFAULT_GOAL := razweb
BUILDFLAGS := -mod=vendor -ldflags="-s -w" -gcflags=-trimpath=$(CURDIR)

razweb:
	go generate $(BUILDFLAGS) ./internal
	go build $(BUILDFLAGS) ./cmd/razweb
