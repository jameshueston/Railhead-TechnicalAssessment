#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE workers;
	GRANT ALL PRIVILEGES ON DATABASE workers TO postgres;
EOSQL

# Load the 'workers' database with tables, constraints and seed records
# Note: 'initsql' is NOT 'init.sql' to prevent auto-running on default 'postgres' db
# https://github.com/docker-library/docs/tree/master/postgres#initialization-scripts
psql --username "$POSTGRES_USER" --dbname=workers --file=/docker-entrypoint-initdb.d/initsql