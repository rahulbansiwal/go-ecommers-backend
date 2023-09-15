DB_URL=postgresql://rahul:admin@localhost:5432/ecom?sslmode=disable
name=ecom
pwd=F:\Code\go backend

schema:
	dbml2sql --postgres -o db/schema/schema.sql db/schema/schema.dbml
postgres-server:
	docker run --name db -e POSTGRES_USER=rahul -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=ecom -p 5432:5432 -d ab8fb914369e
dropdb:
	docker exec -it db dropdb --username=rahul ecom
createdb:
	docker exec -it db createdb --username=rahul --owner=rahul ecom
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
sqlc:
	docker run --rm -v "$(pwd)":/src -w /src sqlc/sqlc generate

.PHONY: schema postgres-server dropdb createdb new_migration migrateup migratedown sqlc
