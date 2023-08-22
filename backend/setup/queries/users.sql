-- Table: public.users

CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255) NOT NULL,
                       password_hash VARCHAR(60) NOT NULL,
                       display_name VARCHAR(255) NOT NULL,
                       is_admin BOOLEAN DEFAULT FALSE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       UNIQUE (email)
);
