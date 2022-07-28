docker-up:
	docker-compose up -d --build --remove-orphans

docker-down:
	docker-compose down

db-migrate-up:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/kajian?sslmode=disable" -verbose up

db-migrate-down:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/kajian?sslmode=disable" -verbose down

server:
	go run main.go

test:
	go test -v -cover ./...

sqlc:
	sqlc generate