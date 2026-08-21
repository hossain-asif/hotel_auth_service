-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_roles
    ADD CONSTRAINT fk_user_roles_user_id
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_roles
    DROP CONSTRAINT IF EXISTS fk_user_roles_user_id;
-- +goose StatementEnd
