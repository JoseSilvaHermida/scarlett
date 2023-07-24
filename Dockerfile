# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Build the Go application inside the container
RUN go build cmd/main.go

# Expose the Unix domain socket as a volume
VOLUME /tmp

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Run the Go application when the container starts
CMD ["./main"]