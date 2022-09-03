# syntax=docker/dockerfile:1.2

## Build
FROM golang:1.19 AS build
WORKDIR /go/src/app
COPY . .
RUN apt update -y && apt upgrade -y && \
    go mod download && \
    mkdir -p /mqtt_enhancer
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/mqtt_enhancer cmd/main.go

## Deploy
FROM gcr.io/distroless/static-debian11
COPY --from=build --chown=nonroot:nonroot /go/bin/mqtt_enhancer /
ENTRYPOINT ["/mqtt_enhancer"]

USER nobody:nobody

CMD ["--help"]

