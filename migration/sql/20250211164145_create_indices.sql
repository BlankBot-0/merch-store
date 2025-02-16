-- +goose Up
-- +goose StatementBegin
create index transactions_from_user_id_idx on transactions (from_user_id);
create index transactions_to_user_id_idx on transactions (to_user_id);
create index purchases_user_id_idx on purchases (user_id);
create unique index users_login_idx on users (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index transactions_from_user_id_idx;
drop index transactions_to_user_id_idx;
drop index purchases_user_id_idx;
drop index users_login_idx;
-- +goose StatementEnd
