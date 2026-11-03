-- +goose Up
CREATE TABLE categories (
  id VARCHAR(50) PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT
);


-- +goose Down
DROP TABLE IF EXISTS categories;
