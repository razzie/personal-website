VERSION ?= `git describe --tags`
BUILDFLAGS := -mod=vendor -ldflags="-s -w" -gcflags=-trimpath=$(CURDIR)
IMAGE_NAME := personal-website
IMAGE_REGISTRY ?= ghcr.io/razzie
FULL_IMAGE_NAME := $(IMAGE_REGISTRY)/$(IMAGE_NAME):$(VERSION)

.PHONY: build
build:
	go build $(BUILDFLAGS) .
	./personal-website build

.PHONY: run
run: build
	./personal-website serve

.PHONY: docker-build
docker-build:
	docker build . -t $(FULL_IMAGE_NAME)

.PHONY: docker-push
docker-push: docker-build
	docker push $(FULL_IMAGE_NAME)
