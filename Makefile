# Simple Makefile for a Go project

# Default target
.DEFAULT_GOAL := all

# Build, setup, and run everything
all: docker-run-detached wait-for-db db-up run

# Start Docker Compose in detached mode so make can continue
docker-run-detached:
	@echo "Starting containers in detached mode..."
	@docker compose up -d --build

# Wait a short period for the DB to initialize (adjust as needed)
wait-for-db:
	@echo "Waiting for database to become available..."
	@sleep 3

# Run the application
run:
	@go run cmd/main.go

db-up:
	@go run cmd/migrate/main.go -op up

db-down:
	@go run cmd/migrate/main.go -op down

# Create DB container (foreground - use when you want to see logs)
docker-run:
	@docker compose up -d --build

# Shutdown DB container
docker-down:
	@docker compose down

clean:
	@echo "Stopping containers and cleaning build artifacts..."
	@docker compose down
	@go clean
	@echo "Clean complete."

.PHONY: all run db-up db-down docker-run docker-run-detached docker-down wait-for-db
