.PHONY: razweb clean

.DEFAULT_GOAL := razweb

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
