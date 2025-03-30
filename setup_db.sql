CREATE DATABASE marchandise;

CREATE USER merchant WITH ENCRYPTED PASSWORD 'e02960b5e1019acb079a1ae860b27c83';

GRANT CONNECT ON DATABASE marchandise TO merchant;

CREATE OR REPLACE FUNCTION created_at_trigger()
RETURNS TRIGGER AS $$
BEGIN
	NEW.created_at := OLD.created_at;
	RETURN NEW;
END:
$$ language plpgsql;

CREATE SCHEMA chat;

CREATE TABLE chat.chat
(
	id text NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

GRANT USAGE ON SCHEMA chat TO merchant;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA chat TO merchant;