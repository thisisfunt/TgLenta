CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    text TEXT UNIQUE NOT NULL,
    ref VARCHAR(400),
    publication_date DATE DEFAULT now(),
    views BIGINT,
    redirected INT,
    channel_id BIGINT REFERENCES channels
)