-- +migrate Up
-- +migrate StatementBegin

ALTER TABLE users ALTER role SET DEFAULT 2;

-- +migrate StatementEnd