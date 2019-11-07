.PHONY: razweb clean
.DEFAULT_GOAL := razweb
BINDATA_TOOL := go run ./tools/go-bindata/

razweb:
	$(BINDATA_TOOL) -prefix assets -o data/bindata.go -pkg data assets/...
	go build ./cmd/razweb

clean:
	go clean
