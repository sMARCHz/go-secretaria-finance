include secret.env

up:
	docker-compose up -d

down:
	docker-compose down

stop:
	docker-compose stop

migrateup:
	migrate -path migrations -database "postgres://smarchz:${DB_PASSWORD}@localhost:5432/secretaria?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgres://smarchz:${DB_PASSWORD}@localhost:5432/secretaria?sslmode=disable" -verbose down

protoc:
	protoc internal/adapters/driving/grpc/proto/finance.proto --go_out=internal/adapters/driving/grpc --go-grpc_out=internal/adapters/driving/grpc

.PHONY: up down stop migrateup migratedown protoc