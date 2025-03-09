CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY,
    tag varchar(16) UNIQUE NOT NULL,
    name varchar(32) UNIQUE NOT NULL
)