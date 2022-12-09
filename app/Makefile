.PHONY: build run clean image

SHELL := /bin/zsh
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

image: build
	-docker rmi ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest -f ${DOCKERFILE} .

# Run the app in a container, connecting to postgres running on the host
rundocker: image
	docker run \
	--env DB_HOST=host.docker.internal --env DB_PORT=5432 --env DB_USER=postgres --env DB_PASSWORD=postgres --env DB_NAME=postgres \
	--publish 8080:8080 \
	--add-host=host.docker.internal:host-gateway \
	${DOCKER_USERNAME}/${APPLICATION_NAME}:latest
	
push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}:latest