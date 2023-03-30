-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE customers (
                                 id SERIAL PRIMARY KEY,
                                 name text not null,
                                 user_id int4 not null,
                                 is_vendor int4 not null,
                                 foreign key (user_id) references users(id)
)

-- +migrate StatementEnd