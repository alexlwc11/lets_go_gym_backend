# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.22 AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /lets-go-gym-backend

# Test stage
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Release stage
FROM scratch as run-release-stage

WORKDIR /app

COPY --from=build-stage /lets-go-gym-backend /lets-go-gym-backend

EXPOSE 8080

CMD [ "/lets-go-gym-backend" ]

# TODO fix the path issue of the private_config.yaml
