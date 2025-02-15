-- +goose Up
-- +goose StatementBegin
create index from_user_id_idx on transactions (from_user_id)
create index to_user_id_idx on transactions (to_user_id)
create index user_id_idx on purchases (user_id)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index from_user_id_idx;
drop index to_user_id_idx;
drop index user_id_idx;
-- +goose StatementEnd
