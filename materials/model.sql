CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(127) UNIQUE,
    order_info jsonb
);