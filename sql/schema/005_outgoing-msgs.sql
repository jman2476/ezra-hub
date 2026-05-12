-- +goose Up
CREATE TABLE outgoingmsgs(
    id UUID PRIMARY,
    encoded_log BYTEA,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    sent BOOLEAN DEFAULT false
);


-- +goose Down
DROP TABLE outgoinglogs;