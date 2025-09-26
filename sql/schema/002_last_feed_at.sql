-- +goose Up
ALTER TABLE feeds
ADD COLUMN last_feed_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN last_feed_at;