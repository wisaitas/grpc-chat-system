-- name: CreateUser :one
INSERT INTO tbl_users (
    id, email, password, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM tbl_users 
WHERE email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM tbl_users 
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM tbl_users
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE tbl_users 
SET email = $2, password = $3, updated_at = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM tbl_users 
WHERE id = $1;
