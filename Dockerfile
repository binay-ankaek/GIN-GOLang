# Stage 1: Build the Go app
FROM golang:1.22.5 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/helloapp/main.go

# Stage 2: Run the Go app
FROM debian:bookworm-slim

# Set working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file (if required)
COPY .env .env

# Expose the port the app will run on
EXPOSE 3000

# Run the binary
CMD ["./main"]

