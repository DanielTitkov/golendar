FROM golang:1.12-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /testing
COPY . . 
COPY ./config.dev.yaml ./config.yaml
COPY ./logger.json .

CMD CGO_ENABLED=0 go test 