-- +goose Up
-- +goose StatementBegin
ALTER TABLE role_permissions
    ADD CONSTRAINT fk_role_permissions_permission_id
    FOREIGN KEY (permission_id)
    REFERENCES permissions(id)
    ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE role_permissions
    DROP CONSTRAINT IF EXISTS fk_role_permissions_permission_id;
-- +goose StatementEnd
