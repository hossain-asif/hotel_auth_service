-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_permissions_name ON permissions (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_permissions_name;
-- +goose StatementEnd
