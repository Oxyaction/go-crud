#!/bin/bash
set -e

# create database account if not exists
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'crud'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE crud"
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'crud_test'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE crud_test"

psql -v ON_ERROR_STOP=1 --username postgres --dbname crud <<-EOSQL
    CREATE EXTENSION "uuid-ossp";
EOSQL

psql -v ON_ERROR_STOP=1 --username postgres --dbname crud_test <<-EOSQL
    CREATE EXTENSION "uuid-ossp";
EOSQL
