# Start from a minimal Go image
FROM golang:1.20-alpine AS build

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and cache the Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

# expose 3000
EXPOSE 3000

# Build the Go application
RUN go build -o app ./cmd/server

# Run the Go application
CMD ["./app"]