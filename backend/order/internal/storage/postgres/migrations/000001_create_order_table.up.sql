CREATE TYPE valid_status AS ENUM ('DRAFT', 'INPROGRESS', 'PAID');
CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status VALID_STATUS
)