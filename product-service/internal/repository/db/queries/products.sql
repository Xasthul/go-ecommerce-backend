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
    sqlc.arg('description')::text,
    sqlc.arg('price_cents'),
    COALESCE(NULLIF(sqlc.arg('currency'), '')::char(3), 'EUR'),
    COALESCE(sqlc.arg('stock')::int4, 0)
);

-- name: GetProductByID :one
SELECT * 
FROM products 
WHERE id = $1;

-- name: GetProducts :many
SELECT * 
FROM products;
