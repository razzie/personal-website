.PHONY: razweb clean
.DEFAULT_GOAL := razweb
BINDATA_TOOL := go run github.com/jteeuwen/go-bindata/go-bindata/

ifeq ($(OS),Windows_NT)
BINDATA_TOOL = tools/windows/go-bindata.exe
else
BINDATA_TOOL = tools/linux/go-bindata
endif

razweb:
	$(BINDATA_TOOL) -prefix assets -o data/bindata.go -pkg data assets/...
	go build ./cmd/razweb

clean:
	go clean
