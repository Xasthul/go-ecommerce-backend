-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (
    user_id,
    token_hash,
    expires_at
) VALUES (
    $1,
    $2,
    $3
);

-- name: GetRefreshToken :one
SELECT *
FROM refresh_tokens
WHERE token_hash = $1
LIMIT 1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token_hash = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM refresh_tokens
WHERE expires_at < NOW();
