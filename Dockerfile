# syntax=docker/dockerfile:1

# Build.
FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app

# Release.
FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /app /app

EXPOSE 8080

USER 1000:1000

ENTRYPOINT ["/app/app"]
