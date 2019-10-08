# get dependencies
FROM golang:1.13 AS dependencies
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

# build stage
FROM dependencies AS build-stage
WORKDIR /app
COPY . . 
RUN cd ./cmd/daemons/sender/ && CGO_ENABLED=0 GOOS=linux go build -o ns

# production stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-stage /app/cmd/daemons/sender/ns .
COPY ./config.dev.yaml ./config.yaml
COPY ./logger.json .
CMD ["./ns"]