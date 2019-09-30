# build stage
FROM golang:latest AS build-stage
WORKDIR /app
COPY . . 
RUN cd ./cmd/daemons/creator/ && CGO_ENABLED=0 GOOS=linux go build -o nc

# production stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-stage /app/cmd/daemons/creator/nc .
COPY ./config.dev.yaml ./config.yaml
COPY ./logger.json .
CMD ["./nc"]