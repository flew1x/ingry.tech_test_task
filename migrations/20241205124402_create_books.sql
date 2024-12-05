-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id uuid PRIMARY KEY,
    title text NOT NULL,
    author text NOT NULL,
    year smallint NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd