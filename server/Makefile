mysql:
	docker run --name mysql_for_cramee -p 3306:3306 -e MYSQL_USER=root -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE="cramee" -d mysql

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

dropdb:
	docker exec -it mysql dropdb 

migrateup:
	migrate -path db/migration -database "mysql://root:secret@localhost:3306/cramee?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "mysql://root:secret@localhost:3306/cramee?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "mysql://root:secret@localhost:3306/cramee?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "mysql://root:secret@localhost:3306/cramee?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...

run:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  cramee/api/repository Store
.PHONY: dropdb migrateup migratedown migrateup1 migratedown1 test server mock run mock createmigrate
