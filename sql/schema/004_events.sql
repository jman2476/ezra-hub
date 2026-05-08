-- +goose Up
CREATE TABLE events(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    owner_id UUID NOT NULL
            references users(id)
            on delete cascade,
    type TEXT NOT NULL,
    occurs_on TIMESTAMP NOT NULL,
    expires_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE events;