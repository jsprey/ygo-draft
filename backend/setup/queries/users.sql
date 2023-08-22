-- Table: public.users

CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL,
                       password_hash VARCHAR(60) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       UNIQUE (username),
                       UNIQUE (email)
);
