CREATE TABLE orders (
    id          UUID PRIMARY KEY,
    user_id     UUID NOT NULL,
    status      TEXT NOT NULL CHECK (status IN ('pending', 'paid')),
    total_cents INTEGER NOT NULL,
    created_at  timestamptz DEFAULT now()
);

CREATE TABLE order_items (
    order_id  UUID REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID NOT NULL,
    quantity  INTEGER NOT NULL,
    price_cents INTEGER NOT NULL,
    PRIMARY KEY (order_id, product_id)
);