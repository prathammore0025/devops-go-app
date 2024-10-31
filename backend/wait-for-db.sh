#!/bin/sh
set -e

until pg_isready -h "$DB_HOST" -U "$DB_USER"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec "$@"

