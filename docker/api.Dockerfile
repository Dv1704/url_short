# Dockerfile
FROM golang:1.23-alpine

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY . .

# Ensure .env is in the image
COPY .env .env

# Build the Go app
RUN go build -o main ./cmd/api/main.go

# Expose the app port
EXPOSE 8080

# Run the app
CMD ["./main"]
