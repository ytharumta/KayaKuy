-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE journal_entries (
    id SERIAL PRIMARY KEY,
    code text not null,
    customer_id int4 not null,
    account_id int4 not null,
    value float not null,
    note text not null,
    user_id int4 not null,
    foreign key (account_id) references accounts(id),
    foreign key (customer_id) references customers(id),
    foreign key (user_id) references users(id)
)

-- +migrate StatementEnd