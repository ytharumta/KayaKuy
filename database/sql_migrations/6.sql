-- +migrate Up
-- +migrate StatementBegin

alter table accounts drop column category_id;

-- +migrate StatementEnd