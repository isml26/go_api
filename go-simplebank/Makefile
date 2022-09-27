postgres:
	docker run --name bank-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres
createdb:
	docker exec -it bank-postgres createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it bank-postgres dropdb --username=postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate	
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc