-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(191) NOT NULL,
    role VARCHAR(16) NOT NULL DEFAULT 'user',
    status BOOLEAN NOT NULL DEFAULT TRUE ,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE users;
