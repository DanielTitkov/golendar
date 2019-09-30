# build stage
FROM golang:latest AS build-stage
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