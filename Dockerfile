# Specify Go image
FROM golang:latest as goapi

# Set the working directory
WORKDIR /app

# Copy the Go application code to the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/server/main.go

# Expose the port your Go application listens on
EXPOSE 8080

# Start the Go application
CMD ["./main"]
