-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_roles_deleted_at ON roles (deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_roles_deleted_at;
-- +goose StatementEnd
