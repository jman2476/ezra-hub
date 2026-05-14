-- +goose Up
CREATE TYPE genre AS ENUM 
        ('ride', 'shopping', 'check-in', 'meal', 'other');
ALTER TABLE events
ADD column min_volunteers INT,
ADD column max_volunteers INT,
ADD column respondants UUID[],
ADD column description TEXT NOT NULL,
DROP column type,
ADD column category GENRE NOT NULL;

-- +goose Down
ALTER TABLE events
DROP column min_volunteers,
DROP COLUMN max_volunteers,
DROP COLUMN respondants,
DROP COLUMN description,
DROP COLUMN category,
ADD COLUMN type TEXT NOT NULL;
DROP TYPE genre;