.DEFAULT_GOAL := run

include .env

build:
	go build -o todo_server cmd/main.go

run: build
	./todo_server

migration:
	cd ./internal/database/migrations && goose postgres postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable up

migration_down:
	cd ./internal/database/migrations && goose postgres postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable down

sqlc:
	sqlc generate
