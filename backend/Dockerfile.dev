FROM golang:1.21-alpine

WORKDIR /app

# Install reflex for hot reloading
RUN go install github.com/cespare/reflex@latest

# Copy modules
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Run with reflex
CMD ["reflex", "-r", "\\.go$", "-s", "--", "sh", "-c", "go run cmd/api/main.go"]
