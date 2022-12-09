.PHONY: build run clean image

SHELL := /bin/zsh
GIT_HASH ?= $(shell git log --format="%h" -n 1)
BINARY_DIR = bin
GOBINARY = rh-workers-app
APPLICATION_NAME = rh-workers-app
DOCKER_USERNAME = jameshueston
DOCKERFILE = Dockerfile

all: fmt clean build deps

#####################################
# GO targets

$(BINARY_DIR):
	@mkdir -p $@

fmt:
	go fmt

clean:
	go clean
	rm -rf ${BINARY_DIR}
	
build: clean
	go build -o ${BINARY_DIR}/${GOBINARY} ./...

deps:
	go mod tidy

runlocal:
	source env.sh && ./${BINARY_DIR}/${GOBINARY}

#####################################
# Docker targets

image:
	-docker rmi ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH}
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH} -f ${DOCKERFILE} .

rundocker: image
	docker run --publish 8080:8080 ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH}

push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH}

release:
	docker pull ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH}
	docker tag  ${DOCKER_USERNAME}/${APPLICATION_NAME}:${GIT_HASH} ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest