#!/bin/bash
set -e

username=${1?"Usage: user postgres_db"}
postgres_db=${2?"Usage: user postgres_db"}

printf "Setup with params username=%s postgres_db=%s\n" $username $postgres_db

psql --username "$username" --dbname "$postgres_db" <<-EOSQL
CREATE DATABASE IF NOT EXISTS marchandise;

CREATE USER merchant WITH ENCRYPTED PASSWORD 'e02960b5e1019acb079a1ae860b27c83';

GRANT CONNECT ON DATABASE marchandise TO merchant;
EOSQL

psql --username "$username" --dbname "marchandise" <<-EOSQL
CREATE OR REPLACE FUNCTION created_at_trigger()
RETURNS TRIGGER AS \$\$
BEGIN
	NEW.created_at := OLD.created_at;
	RETURN NEW;
END
\$\$ language plpgsql;
EOSQL