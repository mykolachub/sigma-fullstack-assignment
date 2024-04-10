CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR (255) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL,
    role VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL
)