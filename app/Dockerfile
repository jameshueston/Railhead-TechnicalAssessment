# syntax=docker/dockerfile:1

# Alpine for small footprint, compared to Ubuntu
FROM golang:1.19-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /rh-workers-app

EXPOSE 8080

CMD [ "/rh-workers-app" ]