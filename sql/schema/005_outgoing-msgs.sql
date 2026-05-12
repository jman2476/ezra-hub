-- +goose Up
CREATE TABLE outgoingmessages(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    encoded_msg BYTEA NOT NULL,
    sent BOOLEAN NOT NULL DEFAULT false 
);


-- +goose Down
DROP TABLE outgoingmessages;