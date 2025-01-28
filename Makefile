.PHONY: up down build test run migrate help

up:
    docker-compose up -d --build

down:
    docker-compose down

build:
	go build -o bin/server backend/main.go

test:
	go test -v ./...

run:
	go run backend/main.go

migrate:
	docker exec -it your-postgres-container /bin/bash -c "psql -U your-db-user -d your-db-name -f /docker-entrypoint-initdb.d/init.sql"

help:
	@echo "make up          - Run docker compose up"
	@echo "make down        - Run docker compose down"
	@echo "make build      - Build backend application"
	@echo "make test       - Run all tests"
	@echo "make run        - Run backend application"
    @echo "make migrate     - Run database migrations"