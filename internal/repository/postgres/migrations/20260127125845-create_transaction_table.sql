-- +migrate Up
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    amount DOUBLE PRECISION,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE transactions;
