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


-- name: GetAllProvider :many
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
LIMIT $1 OFFSET $2;
