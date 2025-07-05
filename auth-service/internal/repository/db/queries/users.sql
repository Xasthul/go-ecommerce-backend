-- name: CreateUser :exec
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
);

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;