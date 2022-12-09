# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /rh-workers-app ./...

## Deploy
# For lean deployments of static binaries: https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10 AS dev

WORKDIR /

COPY --from=build /rh-workers-app /rh-workers-app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/rh-workers-app"]