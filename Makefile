PHONY: build test push

DOCKER_USER ?= duologic
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

pushrm:
	docker run --rm -t \
		-v ${PWD}:/p \
		-e DOCKER_USER='${DOCKER_USER}' \
		-e DOCKER_PASS='${DOCKER_PASS}' \
		chko/docker-pushrm:1 \
		--file /p/README.md \
		--short "gitwatch" \
		${IMAGE_PREFIX}/${IMAGE_NAME}
