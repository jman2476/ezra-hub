-- +goose Up
CREATE TYPE subscription AS ENUM 
        ('ride', 'shopping', 'check-in', 'meal', 'gathering', 'other');
ALTER TABLE users
ADD COLUMN subs subscription[];

-- +goose Down
ALTER TABLE users
DROP COLUMN subs;
DROP TYPE subscription;