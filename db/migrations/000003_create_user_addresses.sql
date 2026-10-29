-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE user_addresses (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    label        TEXT,                          -- 'Home', 'Office'
    address_line TEXT NOT NULL,                 -- tole / street / ward details
    city         TEXT NOT NULL,                 -- municipality / city
    location     geography(Point, 4326),        -- ST_MakePoint(longitude, latitude)
    is_default   BOOLEAN NOT NULL DEFAULT false,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX user_addresses_user_id_idx
    ON user_addresses (user_id);

CREATE INDEX user_addresses_location_idx
    ON user_addresses USING GIST (location);

-- At most one default address per user
CREATE UNIQUE INDEX user_addresses_one_default_idx
    ON user_addresses (user_id)
    WHERE is_default;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_addresses;

-- The postgis extension is intentionally not dropped: other tables
-- (e.g. providers.location) may depend on it.
-- DROP EXTENSION IF EXISTS postgis;
-- +goose StatementEnd
