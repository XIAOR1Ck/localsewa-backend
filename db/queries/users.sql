-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  password,
  phone,
  role
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
  )
RETURNING *;

-- name: GetUserByID :one
SELECT id, email, name, role, created_at, updated_at
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, email, name, role, created_at, updated_at
FROM users
WHERE email = $1
LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET name = $2,
    email = $3,
    role = $4,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, email, name, role, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
