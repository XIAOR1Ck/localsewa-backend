-- +goose Up
CREATE TABLE providers (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name           VARCHAR(255) NOT NULL,
    email          VARCHAR(255) NOT NULL UNIQUE,
    phone          VARCHAR(15) NOT NULL UNIQUE,
    password_hash  TEXT NOT NULL,
    profile_image  TEXT,
    category       TEXT,
    experience     REAL NOT NULL DEFAULT 0,
    hourly_rate    REAL,
    description    TEXT,
    rating         REAL NOT NULL DEFAULT 0,
    total_reviews  INT  NOT NULL DEFAULT 0,
    created_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS providers;

