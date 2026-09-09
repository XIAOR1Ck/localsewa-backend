-- name: CreateUser :one
INSERT INTO users (
  id,
  name,
  email,
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

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;
