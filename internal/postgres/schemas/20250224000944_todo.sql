-- +goose Up
-- +goose StatementBegin
CREATE TABLE todos (
    id          SERIAL  PRIMARY KEY,
    title       TEXT    NOT NULL,
    description TEXT    NOT NULL,
    completed   BOOLEAN NOT NULL DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE todos;
-- +goose StatementEnd
