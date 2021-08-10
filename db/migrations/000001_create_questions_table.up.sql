CREATE TABLE IF NOT EXISTS questions (
    id serial PRIMARY KEY,
    type VARCHAR(30),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    content TEXT NOT NULL UNIQUE
    );