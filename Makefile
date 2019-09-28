.PHONY: bindata razweb clean

.DEFAULT_GOAL := razweb

bindata:
	go run github.com/jteeuwen/go-bindata/go-bindata/ -prefix assets assets/...

ifeq ($(OS),Windows_NT)
razweb: bindata
	go build -o bin/razweb.exe
else
razweb: bindata
	go build -o bin/razweb
endif

clean:
	go clean
