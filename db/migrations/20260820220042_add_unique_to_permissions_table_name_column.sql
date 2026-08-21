-- +goose Up
-- +goose StatementBegin
ALTER TABLE permissions
ADD CONSTRAINT uq_permissions_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE permissions
DROP CONSTRAINT IF EXISTS uq_permissions_name;
-- +goose StatementEnd