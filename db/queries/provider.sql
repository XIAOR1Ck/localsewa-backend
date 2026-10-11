-- name: CreateProvider :one
INSERT INTO providers (
  name,
  email,
  password,
  phone
)
VALUES (
  $1,
  $2,
  $3,
  $4
  )
RETURNING *;

-- name: GetProviderByID :one
SELECT id, email, name, created_at, updated_at
FROM providers
WHERE id = $1
LIMIT 1;

-- name: GetProviderByEmail :one
SELECT id, email, name, created_at, updated_at
FROM providers
WHERE email = $1
LIMIT 1;

-- name: UpdateProvider :exec
UPDATE providers
SET name = $2,
    email = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteProvider :exec
DELETE FROM providers
WHERE id = $1;

-- name: ListProvider :many
SELECT id, email, name, created_at, updated_at
FROM providers
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
