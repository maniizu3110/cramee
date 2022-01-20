#!/bin/sh

# non-zeroが返されたときにreturnすることを指定
set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

# takes all parameters passed to the script and run it
echo "start the app"
exec "$@"
