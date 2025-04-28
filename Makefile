include .env
export $(shell sed 's/=.*//' .env)

# Khởi chạy PostgreSQL với Docker
postgres:
	docker run --name $(POSTGRES_CONTAINER) -p $(DB_PORT):5432 \
		-e POSTGRES_USER=$(DB_USER) -e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-d postgres:latest

# Tạo database bên trong container
createdb:
	docker exec -it $(POSTGRES_CONTAINER) createdb --username=$(DB_USER) --owner=$(DB_USER) $(DB_NAME)

# Xóa container PostgreSQL
stopdb:
	docker stop $(POSTGRES_CONTAINER) && docker rm $(POSTGRES_CONTAINER)

# Chạy migration (giả sử bạn dùng migrate)
migrateup:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

add-migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)
	
sqlc:
	sqlc generate

run: 
	go run main.go