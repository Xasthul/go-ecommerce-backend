-- name: CreateProduct :exec
INSERT INTO products (
    id, 
    category_id,
    name,
    description,
    price_cents,
    currency,
    stock
) VALUES (
    gen_random_uuid(),  
    sqlc.arg('category_id')::int2,
    sqlc.arg('name'),
    sqlc.narg('description')::text,
    sqlc.arg('price_cents'),
    COALESCE(sqlc.narg('currency')::char(3), 'EUR'),
    COALESCE(sqlc.narg('stock')::int4, 0)
);

-- name: GetProductByID :one
SELECT * 
FROM products 
WHERE id = $1;

-- name: GetProducts :many
SELECT * 
FROM products;

-- name: UpdateProduct :one
UPDATE products
SET
    category_id = COALESCE(sqlc.narg('category_id')::int2,   category_id),
    name        = COALESCE(sqlc.narg('name'),                name),
    description = COALESCE(sqlc.narg('description')::text,   description),
    price_cents = COALESCE(sqlc.narg('price_cents'),         price_cents),
    currency    = COALESCE(
                     sqlc.narg('currency')::char(3),
                     currency
                 ),
    stock       = COALESCE(sqlc.narg('stock')::int4,         stock),
    updated_at = now()
WHERE id = sqlc.arg('id')::uuid
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = sqlc.arg('id')::uuid;

-- name: DecreaseStock :one
UPDATE products
SET stock = stock - $2
WHERE id = $1 AND stock >= $2
RETURNING *;
