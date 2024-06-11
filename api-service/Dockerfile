# syntax=docker/dockerfile:1

FROM golang:alpine as build

WORKDIR /go/src/github.com/will-kerwin/go-microservice-bookstore
ARG GIT_TAG_NAME=1.0.0

ENV CGO_ENABLED 1
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk --no-cache update && \
    apk --no-cache add git gcc libc-dev

COPY ../go.mod ./
COPY ../go.sum ./

RUN go mod download
RUN go mod verify

COPY ./api-service ./api-service
COPY ../pkg ./pkg
COPY ../gen ./gen
COPY ../internal ./internal
COPY ../docs ./docs

RUN go build -tags musl -ldflags "-s -w -X main.Version=$GIT_TAG_NAME" -o /bin/service ./api-service/cmd

FROM alpine:latest as final
WORKDIR /app
ENV PORT=8080

COPY --from=build /bin/service /service

EXPOSE ${PORT}

ENTRYPOINT [ "/service" ]