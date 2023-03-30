-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    full_name varchar(256),
    user_name varchar(256),
    email varchar(256),
    avatar varchar(256),
    social_id varchar(256),
    provider varchar(256),
    role int4
);

-- +migrate StatementEnd