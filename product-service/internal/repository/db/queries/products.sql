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
    gen_random_uuid(),  -- id
    $1,                  -- category_id
    $2,                  -- name
    $3,                  -- description
    $4,                  -- price_cents
    COALESCE($5, 'EUR'), -- currency (default to EUR)
    COALESCE($6, 0)      -- stock (default to 0)
);

-- name: GetProductByID :one
SELECT * 
FROM products 
WHERE id = $1;

-- name: GetProducts :many
SELECT * 
FROM products;
