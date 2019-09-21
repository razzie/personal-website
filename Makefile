.PHONY: all deps razweb clean

.DEFAULT_GOAL := all

all: deps razweb

deps:
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...
	go get github.com/google/go-github/...
	go get golang.org/x/oauth2/...

ifeq ($(OS),Windows_NT)
razweb:
	go-bindata-assetfs.exe -prefix assets assets/...
	go build -o bin/razweb.exe
else
razweb:
	go-bindata-assetfs -prefix assets assets/...
	go build -o bin/razweb
endif

clean:
	go clean
