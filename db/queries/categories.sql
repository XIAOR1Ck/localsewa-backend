-- name: GetAllCategories :many
SELECT * FROM categories;


-- name: AddCategory :one
INSERT INTO categories (
    id,
    name,
    description
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;
