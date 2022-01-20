mysql:
	docker run --name mysql_for_cramee -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE="cramee" -d mysql

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

dropdb:
	docker exec -it mysql dropdb 

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/cramee" -verbose up

migrateup1:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/cramee" -verbose up 1

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/cramee" -verbose down

migratedown1:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/cramee" -verbose down 1

test:
	go test -v -cover ./...

run:
	go run main.go

clear_db:
	go run db/tools/clear_db/clear_db.go

create-sql: # 指定するデータベースの構造を削除->gormのautomigrateを走らせる->構図をsqlファイルとして取り出す
	go run db/tools/clear_db/clear_db.go
	go run db/tools/migrate/migrate.go
	mysqldump --no-data -h 127.0.0.1 -u root -p cramee > ./dump.sql


mock:
	mockgen -package mockdb -destination db/mock/store.go  cramee/api/repository Store
.PHONY: dropdb migrateup migratedown migrateup1 migratedown1 test server mock run mock createmigrate create-sql
