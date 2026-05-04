-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL, 
    update_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    email TEXT NOT NULL,
    UNIQUE (name, phone_number),
    UNIQUE (name, email)
);

-- +goose Down
DROP TABLE users;