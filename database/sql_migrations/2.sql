-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE users add column password text;

-- +migrate StatementEnd