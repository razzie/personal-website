build:
	go generate ./cmd/razweb
	go build -mod=vendor ./cmd/razweb

.PHONY: build