-- name: CreateUser :exec
INSERT INTO users (
    id,
    email,
    password_hash,
    role,
    status
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    COALESCE($3, 0),
    0
);

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;