#!/bin/sh

# non-zeroが返されたときにreturnすることを指定
set -e

echo "run db migration"
source /app/app.env
# dockerの作成時だけはホストをlocalhostからpostgresに変更する
export DB_SOURCE="postgres://root:secret@postgres:5432/crameee?sslmode=disable"
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

# takes all parameters passed to the script and run it
echo "start the app"
exec "$@"
