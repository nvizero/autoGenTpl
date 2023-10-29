postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=a123456 -d postgres:12-alpine

droppq:
	docker rm -f postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root gen_laravel

dropdb:
	docker exec -it postgres dropdb gen_laravel 

migrateup:
	migrate -path db/migration -database "postgresql://root:a123456@localhost:5432/gen_laravel?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:a123456@localhost:5432/gen_laravel?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup sqlc

