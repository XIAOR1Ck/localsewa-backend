-- name: CreateProvider :one
INSERT INTO providers (
  name,
  email,
  password_hash,
  phone
)
VALUES (
  $1,
  $2,
  $3,
  $4
  )
RETURNING *;

-- name: GetProviderForAuth :one
SELECT p.id, p.name, p.email, p.password_hash
FROM providers p
WHERE email = $1
LIMIT 1;


-- name: GetProviderByID :one
SELECT
    id,
    name,
    email,
    phone,
    profile_image,
    category,
    experience,
    hourly_rate,
    description,
    rating,
    total_reviews,
    created_at,
    updated_at
FROM providers
WHERE id = $1;


-- name: GetProviderByEmail :one
SELECT id, email, name, created_at, updated_at
FROM providers
WHERE email = $1;



-- name: DeleteProvider :exec
DELETE FROM providers
WHERE id = $1;




-- name: GetAllProviders :many
SELECT
    id,
    name,
    email,
    phone,
    profile_image,
    category,
    experience,
    hourly_rate,
    description,
    rating,
    total_reviews,
    created_at,
    updated_at
FROM providers
WHERE (sqlc.narg('categories')::text[] IS NULL OR category = ANY(sqlc.narg('categories')::text[]))
  AND (sqlc.narg('min_rate')::real IS NULL OR hourly_rate >= sqlc.narg('min_rate')::real)
  AND (sqlc.narg('max_rate')::real IS NULL OR hourly_rate <= sqlc.narg('max_rate')::real)
  AND (sqlc.narg('min_rating')::real IS NULL OR rating >= sqlc.narg('min_rating')::real)
  AND (sqlc.narg('min_experience')::real IS NULL OR experience >= sqlc.narg('min_experience')::real)
  AND (sqlc.narg('max_experience')::real IS NULL OR experience <= sqlc.narg('max_experience')::real)
  AND (sqlc.narg('search')::text IS NULL OR name ILIKE '%' || sqlc.narg('search')::text || '%')
ORDER BY created_at DESC, id
LIMIT sqlc.arg('limit') OFFSET sqlc.arg('offset');
