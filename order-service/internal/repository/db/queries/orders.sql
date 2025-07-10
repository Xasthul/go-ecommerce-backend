-- name: CreateOrder :one
INSERT INTO orders (
    id,
    user_id,
    status,
    total_cents
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
) 
RETURNING *;

-- name: GetOrderById :one
SELECT *
FROM orders
WHERE id = $1;

-- name: GetUserOrders :many
SELECT *
FROM orders
WHERE user_id = $1;