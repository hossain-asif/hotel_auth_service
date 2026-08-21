-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_roles
    ADD CONSTRAINT fk_user_roles_role_id
    FOREIGN KEY (role_id)
    REFERENCES roles(id)
    ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_roles
    DROP CONSTRAINT IF EXISTS fk_user_roles_role_id;
-- +goose StatementEnd
