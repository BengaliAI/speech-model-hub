FROM golang:1.21.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
ARG ENVIRONMENT
ENV ENVIRONMENT=$ENVIRONMENT
ENV GIN_MODE=release
RUN go mod download
COPY . /app
RUN go build -o speech-model-hub

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ /app/
ENTRYPOINT ./speech-model-hub
