-- +goose Up
-- +goose StatementBegin
create table users
(
    id    bigserial primary key,
    coins integer not null
);

create table transactions
(
    from_user_id integer,
    to_user_id   integer,
    coins        integer,

    done_at timestamp not null default now()
);

create table merchandise
(
    id bigserial primary key,
    name text,
    coins integer
);

create table purchases
(
    user_id integer,
    item_id integer,

    purchased_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table purchases;
drop table merchandise;
drop table transactions;
drop table users;
-- +goose StatementEnd
