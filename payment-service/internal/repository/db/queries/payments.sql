-- name: CreatePayment :exec
INSERT INTO payments (
    id,
    order_id,
    user_id,
    amount_cents,
    status
) VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
);
