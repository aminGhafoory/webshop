-- name: GetAuthor :one
SELECT  *
FROM users
WHERE user_id = $1;