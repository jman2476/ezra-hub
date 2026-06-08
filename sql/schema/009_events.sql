-- +goose Up
ALTER TABLE events
ADD COLUMN location TEXT NOT NULL DEFAULT 'unset';

-- +goose Down
ALTER TABLE events
DROP COLUMN location;