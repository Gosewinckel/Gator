-- +goose Up
CREATE TABLE feed_follows (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	user_id UUID,
	feed_id UUID,
	FOREIGN KEY(user_id)
		REFERENCES users(id)
	FOREIGN KEY(feed_id)
		REFERENCES feed(id)
)
