-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age SMALLINT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS users;
