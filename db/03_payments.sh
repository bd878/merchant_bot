#!/bin/bash
set -e

username=${1?"Usage: user"}

printf "Create with username %s\n" $username

psql -v ON_ERROR_STOP=1 --username "$username" --dbname "marchandise" <<-EOSQL
CREATE SCHEMA IF NOT EXISTS payments;

CREATE TABLE IF NOT EXISTS payments.payments
(
	id bigint NOT NULL,
	user_id integer NOT NULL,
	refunded bool NOT NULL,
	telegram_payment_charge_id text NOT NULL,
	provider_payment_charge_id text NOT NULL,
	invoice_payload text NOT NULL,
	currency text NOT NULL,
	total_amount integer NOT NULL,
	created_at timestamptz NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

CREATE TRIGGER created_at_payments_trgr BEFORE UPDATE ON payments.payments FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();

GRANT USAGE ON SCHEMA payments TO merchant;
GRANT INSERT, UPDATE, DELETE, SELECT ON ALL TABLES IN SCHEMA payments TO merchant;
EOSQL