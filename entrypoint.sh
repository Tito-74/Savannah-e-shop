#!/bin/bash

set -e

host="$1"
port="$2"
shift
cmd="$@"

until psql -h "$host" -p "$port" -U "postgres" -c '\q' &> /dev/null; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd
