.PHONY: razweb clean

.DEFAULT_GOAL := razweb

#bindata:
#	go run github.com/jteeuwen/go-bindata/go-bindata/ -prefix assets assets/...

ifeq ($(OS),Windows_NT)
razweb:
	tools/windows/go-bindata.exe -prefix assets assets/...
	go build -o bin/razweb.exe
else
razweb:
	tools/linux/go-bindata -prefix assets assets/...
	go build -o bin/razweb
endif

clean:
	go clean
