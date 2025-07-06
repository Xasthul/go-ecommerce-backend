CREATE TABLE categories (
    id   SMALLSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE products (
    id          UUID PRIMARY KEY,
    category_id SMALLINT REFERENCES categories(id),
    name        TEXT NOT NULL,
    description TEXT,
    price_cents INTEGER NOT NULL,
    currency    CHAR(3) NOT NULL DEFAULT 'EUR',
    stock       INTEGER NOT NULL DEFAULT 0,
    created_at  timestamptz NOT NULL DEFAULT now(),
    updated_at  timestamptz NOT NULL DEFAULT now()
);
