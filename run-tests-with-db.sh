#!/bin/bash

set -e
set -u

SCRIPTPATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"

POSTGRES_HOST="localhost"
POSTGRES_PORT=5433
POSTGRES_DB_NAME="postgres"
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="postgres"
DRIVER_NAME="postgres"     

# connect with: psql -h localhost -p 5433  -U postgres --password
docker run --name testing-postgres --rm -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD -d -p $POSTGRES_PORT:5432 postgres

echo "waiting for database to start..."

RETRIES=30
until PGPASSWORD=$POSTGRES_PASSWORD psql -h localhost -p $POSTGRES_PORT -U $POSTGRES_USER -c "select 1" > /dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
  echo "Waiting for postgres server, $((RETRIES--)) remaining attempts..."
  sleep 1
done

clean_up () {
    ARG=$?
    echo "> clean_up with exitcode: $ARG"
    docker stop testing-postgres
    exit $ARG
}
trap clean_up EXIT

ACCESS_SECRET="abcdefghijklmnopqrstuvwxyz" \
POSTGRES_HOST=$POSTGRES_HOST \
POSTGRES_PORT=$POSTGRES_PORT  \
POSTGRES_USER=$POSTGRES_USER \
POSTGRES_PASSWORD=$ \
POSTGRES_DATABASE=$POSTGRES_DB_NAME \
DRIVER_NAME=$DRIVER_NAME \
MIGRATION_FILES=./migrations/\
TEST_WITH_DB=true\
TEST_SEED_FILE="./seed/init.sql"\

go test -v ./...

