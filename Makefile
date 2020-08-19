PHONY: build test push

IMAGE_NAME ?= gitwatch
IMAGE_PREFIX ?= duologic
IMAGE_TAG ?= 0.0.1

KUBECONFIG ?= ${HOME}/.kube/config

build:
	docker build -t ${IMAGE_PREFIX}/${IMAGE_NAME} .
	docker tag ${IMAGE_PREFIX}/${IMAGE_NAME} ${IMAGE_PREFIX}/${IMAGE_NAME}:$(IMAGE_TAG)

test:
	go vet ./...
	go build ./...

push:
	docker push ${IMAGE_PREFIX}/${IMAGE_NAME}:$(IMAGE_TAG)
	docker push ${IMAGE_PREFIX}/${IMAGE_NAME}:latest
