-- name: GetUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1 OFFSET $2;
