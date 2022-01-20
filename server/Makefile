postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB="crameee" -d postgres:13-alpine

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

dropdb:
	docker exec -it postgres13 dropdb 

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/crameee?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/crameee?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/crameee?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/crameee?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...

run:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  crameee/api/repository Store
.PHONY: postgres dropdb migrateup migratedown migrateup1 migratedown1 test server mock run mock createmigrate postgres
