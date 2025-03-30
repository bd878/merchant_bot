#!/bin/bash
set -e

username=${1?"Usage: user"}

printf "Create with username %s\n" $username

psql -v ON_ERROR_STOP=1 --username "$username" --dbname "marchandise" <<-EOSQL
CREATE SCHEMA IF NOT EXISTS chat;

CREATE TABLE IF NOT EXISTS chat.chat
(
	id text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

GRANT USAGE ON SCHEMA chat TO merchant;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA chat TO merchant;
EOSQL