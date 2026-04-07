.PHONY: run docker-up migrate

# Configuration
DB_URL="postgres://myuser:mypassword@localhost:5432/myappdb?sslmode=disable"

# Run backend locally
run:
	@echo "Running backend locally..."
	cd backend && go run cmd/api/main.go

# Start Docker containers
docker-up:
	@echo "Starting Docker containers in background..."
	docker-compose --env-file .env.development up -d --build

# Run database migrations
migrate:
	@echo "Running database migrations..."
	migrate -path migrations -database $(DB_URL) up
