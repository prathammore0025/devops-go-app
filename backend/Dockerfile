# Step 1: Use the official Golang image to build the application
FROM golang:1.19 AS builder
WORKDIR /app

# Copy all Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o app main.go

# Step 2: Use a smaller image to run the built application
FROM debian:bookworm-slim
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/app .

# Set environment variables for the database connection (optional)
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./app"]

