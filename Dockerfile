# Dockerfile
# Use the official Golang image as the base image
FROM golang:1.21.0 AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . ./

# Build the application
RUN go build -o monitoring-agent ./monitoring/agent.go

# Start a new stage from scratch
FROM alpine:latest  

# Install necessary dependencies for running the Go binary
RUN apk --no-cache add ca-certificates

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the build stage to the current directory inside the container
COPY --from=build /app/monitoring-agent .

# Run the monitoring agent when the container starts
CMD ["./monitoring-agent"]
