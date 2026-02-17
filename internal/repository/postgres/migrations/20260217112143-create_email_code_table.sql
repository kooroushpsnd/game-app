-- +migrate Up
CREATE TABLE IF NOT EXISTS email_codes (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    hash_code CHAR(64) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    attempts INT NOT NULL DEFAULT 0,
    expiration_date TIMESTAMP WITH TIME ZONE NOT NULL,
    user_id INT NOT NULL,

    CONSTRAINT fk_email_codes_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_email_codes_email ON email_codes(email);
CREATE INDEX IF NOT EXISTS idx_email_codes_user_id ON email_codes(user_id);
CREATE INDEX IF NOT EXISTS idx_email_codes_expiration_date ON email_codes(expiration_date);

-- +migrate Down
DROP TABLE IF EXISTS email_codes;