.PHONY: run build test migrate-up migrate-down sqlc-gen docker-up docker-down

# Development
run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

# Database migrations
migrate-up:
	goose up

migrate-down:
	goose down

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

# sqlc
sqlc-gen:
	sqlc generate

# Docker
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Clean
clean:
	rm -rf bin/
	rm -rf internal/db/
