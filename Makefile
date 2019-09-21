.PHONY: all deps razweb clean

.DEFAULT_GOAL := all

all: deps razweb

deps:
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/elazarl/go-bindata-assetfs/...
	go get github.com/google/go-github/...
	go get golang.org/x/oauth2/...

ifeq ($(OS),Windows_NT)
razweb: bindata.go
	go build -o bin/razweb.exe

bindata.go:
	go-bindata-assetfs.exe -prefix assets assets/...
else
razweb: bindata.go
	go build -o bin/razweb

bindata.go:
	go-bindata-assetfs -prefix assets assets/...
endif

clean:
	go clean
