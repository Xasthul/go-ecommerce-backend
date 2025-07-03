-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
    user_id,
    token_hash,
    expires_at
) VALUES (
    $1,              -- user_id
    $2,              -- sha‑256(token)
    $3               -- expires_at (TIMESTAMPTZ)
)
RETURNING id, user_id, token_hash, expires_at, created_at;

-- name: GetRefreshToken :one
SELECT *
FROM refresh_tokens
WHERE token_hash = $1
LIMIT 1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE token_hash = $1;

-- name: DeleteTokensByUser :exec
DELETE FROM refresh_tokens
WHERE user_id = $1;

-- name: DeleteExpiredTokens :exec
DELETE FROM refresh_tokens
WHERE expires_at < NOW();
