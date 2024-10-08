# Step 1: Build the Go app
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app binary
RUN go build -o main .

# Step 2: Use a minimal Alpine Linux image for production
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app listens on
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
