# Use an official Golang runtime as a parent image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

                                                                                                              
