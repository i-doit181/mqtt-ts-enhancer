# syntax=docker/dockerfile:1.2

ARG goarch

## Build
FROM golang:1.20 AS build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=${goarch:-amd64}
WORKDIR /go/src/app
COPY . .
RUN go mod download && \
    mkdir -p /mqtt_enhancer
RUN go build -o /go/bin/mqtt_enhancer cmd/main.go

## Deploy
FROM gcr.io/distroless/static:nonroot-${goarch:-amd64}
COPY --from=build --chown=nonroot:nonroot /go/bin/mqtt_enhancer /
ENTRYPOINT ["/mqtt_enhancer"]

USER nobody:nobody

CMD ["--help"]

