-- name: CreateOrderItem :exec
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    price_cents
) VALUES (
    $1,
    $2,
    $3,
    $4
);

-- name: GetOrderItemsForOrder :many
SELECT *
FROM order_items
WHERE order_id = $1;