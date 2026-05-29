-- +goose Up
CREATE TABLE posts (
	id UUID PRIMARY NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	title TEXT,
	url TEXT NOT NULL,
	description TEXT,
	published_at TIMESTAMP,
	feed_id UUID NOT NULL,
	FOREIGN KEY(feed_id)
		REFERENCES feeds(id)
);

-- +goose Down
DROP TABLE posts;
