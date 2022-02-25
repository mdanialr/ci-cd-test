# syntax=docker/dockerfile:1

# Use golang version 1.17 with os alpine linux
FROM golang:1.17-alpine

# Change pwd to /app
WORKDIR /app

# Download dependencies
COPY go.mod .
RUN go mod download

# Copy source code to images container
ADD . .

# Compile source code
RUN go build -o bin/main

# Show port 4000 to outside world
EXPOSE 4000

CMD [ "bin/main" ]