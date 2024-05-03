DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'valid_status'
) THEN CREATE TYPE valid_status AS ENUM ('DRAFT', 'INPROGRESS', 'PAID');
END IF;
END $$;
CREATE TABLE IF NOT EXISTS orders (
    order_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status VALID_STATUS
)