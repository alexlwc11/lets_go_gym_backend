# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21 AS build-stage

# Set destination for COPY
WORKDIR /app

#Download Go modules
COPY go.mod go.sum ./
RUN go mod Download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./