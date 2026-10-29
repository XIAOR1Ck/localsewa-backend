
-- +goose Up
-- +goose StatementBegin
CREATE TABLE provider_addresses (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    provider_id  UUID NOT NULL UNIQUE REFERENCES providers(id) ON DELETE CASCADE,
    address_line TEXT NOT NULL,                 -- tole / street / ward details
    city         TEXT NOT NULL,                 -- municipality / city
    location     geography(Point, 4326),        -- ST_MakePoint(longitude, latitude)
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX provider_addresses_location_idx
    ON provider_addresses USING GIST (location);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS provider_addresses;
-- +goose StatementEnd
