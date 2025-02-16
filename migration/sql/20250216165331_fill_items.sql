-- +goose Up
-- +goose StatementBegin
insert into items (type, coins)
values ('t-shirt', 80),
       ('cup', 20),
       ('book', 50),
       ('pen', 10),
       ('powerbank', 200),
       ('hoody', 300),
       ('umbrella', 200),
       ('socks', 10),
       ('wallet', 50),
       ('pink-hoody', 500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
truncate table items;
-- +goose StatementEnd
