# syntax=docker/dockerfile:1

# Use golang version 1.17 with os alpine linux
FROM golang:1.17-alpine

# Change pwd to /app
WORKDIR /app

# Download live reload tools: air
RUN go install github.com/cosmtrek/air@latest

CMD air
