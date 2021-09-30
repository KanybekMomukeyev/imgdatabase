CREATE TABLE IF NOT EXISTS payments (
    payment_id BIGSERIAL PRIMARY KEY NOT NULL,
    total_order_price REAL,
    discount REAL,
    total_price_with_discount REAL,
    minus_price REAL,
    plus_price REAL,
    comment varchar (400)
);