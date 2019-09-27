# build stage
FROM golang:latest AS build-stage
WORKDIR /app
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# production stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-stage /app/main .
COPY ./config.dev.yaml ./config.yaml
COPY ./logger.json .
CMD ["./main"]