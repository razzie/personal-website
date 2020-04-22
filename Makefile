build:
	go generate -mod=vendor ./cmd/razweb
	go build -mod=vendor ./cmd/razweb

.PHONY: build
