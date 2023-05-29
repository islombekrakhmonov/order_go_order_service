CREATE TABLE IF NOT EXISTS orders(
    id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255) NOT NULL,
    totalsum INT NOT NULL
);

CREATE TABLE IF NOT EXISTS products(
    id  VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL
);

INSERT INTO products values(
    '405cb824-fcd1-11ed-be56-0242ac120002', 'Coca Cola', 14000
)
