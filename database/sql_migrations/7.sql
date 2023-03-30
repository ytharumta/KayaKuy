-- +migrate Up
-- +migrate StatementBegin
CREATE TYPE transaction_type AS ENUM ('Debit','Credit');

alter table journal_entries add column transaction_type transaction_type;
alter table journal_entries add column created_at timestamp, add column updated_at timestamp;

-- +migrate StatementEnd