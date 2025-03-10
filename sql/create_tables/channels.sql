CREATE TABLE IF NOT EXISTS channels (
    id BIGSERIAL PRIMARY KEY,
    name varchar(128) UNIQUE NOT NULL,
    tg_id BIGINT UNIQUE NOT NULL,
    subscribers_count BIGINT,
    category_id BIGINT REFERENCES categories,
    logo_ref VARCHAR(400),
    ref VARCHAR(400),
    premium BOOLEAN DEFAULT false
)