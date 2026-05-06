-- +goose Up
CREATE TABLE refreshtokens(
    token TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL
            references users(id)
            on delete cascade,
    expires_at TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP
);

-- +goose Down
DROP TABLE refreshtokens;