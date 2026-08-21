-- +goose Up
-- +goose StatementBegin
ALTER TABLE roles
ADD CONSTRAINT uq_roles_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE roles
DROP CONSTRAINT IF EXISTS uq_roles_name;
-- +goose StatementEnd