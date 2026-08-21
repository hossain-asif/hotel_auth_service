-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_roles_name ON roles (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_roles_name;
-- +goose StatementEnd
