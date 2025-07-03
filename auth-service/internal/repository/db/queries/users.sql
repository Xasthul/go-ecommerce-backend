-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    password_hash,
    role,
    status
) VALUES (
    gen_random_uuid(),
    $1,             -- email
    $2,             -- password_hash
    COALESCE($3, 0),-- role (0 = buyer) ; allow NULL -> default
    0               -- status = pending_email
)
RETURNING *;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC
LIMIT $1               -- page size
OFFSET $2;             -- (page ‑ 1) * size

-- name: UpdateUserStatus :exec
UPDATE users
SET status = $2,       -- new status
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
