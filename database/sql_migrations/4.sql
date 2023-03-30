-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    name varchar(256) not null,
    category_id int4 not null,
    user_id int4 not null,
    foreign key (user_id) references users(id)
)

-- +migrate StatementEnd