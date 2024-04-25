CREATE TABLE IF NOT EXISTS order_items (
    item_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    reserved_id INT NOT NULL,
    FOREIGN KEY(order_id) REFERENCES orders(order_id) ON DELETE NO ACTION
)