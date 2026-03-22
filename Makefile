generate:
    sqlc generate

migrate-up:
    migrate -path db/migrations -database $(DATABASE_URL) up

migrate-down:
    migrate -path db/migrations -database $(DATABASE_URL) down

run:
    go run ./cmd/server