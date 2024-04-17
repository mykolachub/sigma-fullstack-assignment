CREATE TABLE reserved_products (
    reserved_id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    FOREIGN KEY(product_id) REFERENCES products(product_id) ON DELETE NO ACTION
)