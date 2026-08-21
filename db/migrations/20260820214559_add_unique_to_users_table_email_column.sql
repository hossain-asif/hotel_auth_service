-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD CONSTRAINT uq_users_email UNIQUE (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP CONSTRAINT IF EXISTS uq_users_email;
-- +goose StatementEnd