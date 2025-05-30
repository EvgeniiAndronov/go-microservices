-- migrate:up
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       email VARCHAR(100) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       last_login TIMESTAMP,
                       is_active BOOLEAN DEFAULT TRUE
);
-- migrate:down
DROP TABLE users;
