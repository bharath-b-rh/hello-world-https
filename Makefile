export GOOS=linux

IMAGE ?= quay.io/bharath-b-rh/hello-world-https
TAG ?= latest

all: build image push

.PHONY: build
build:
	mkdir -p bin
	go build -o bin/hello-world-https main.go

.PHONY: image
image:
	imagebuilder -t $(IMAGE):$(TAG) .

.PHONY: push
push:
	docker push $(IMAGE):$(TAG)
