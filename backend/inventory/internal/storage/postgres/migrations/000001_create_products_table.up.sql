CREATE TABLE IF NOT EXISTS products (
    product_id SERIAL,
    name VARCHAR (255) NOT NULL,
    price INT NOT NULL,
    quantity INT NOT NULL,
    PRIMARY KEY(product_id)
)