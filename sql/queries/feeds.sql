-- name: GetFeeds :many
SELECT * FROM feeds WHERE user_id = $1;


-- name: CreateFeed :one
INSERT INTO feeds (created_at, updated_at, title, url, user_id)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeedsWithUsername :many
SELECT f.*, u.*
FROM feeds AS f
JOIN users AS u ON f.user_id = u.id;

-- name: MarkFeedAsFetched :exec
UPDATE feeds
SET last_feed_at = NOW(), updated_at = NOW()
WHERE id = $1;


-- name: GetNextFeedToFetch :one
SELECT id, url
FROM feeds
ORDER BY last_feed_at NULLS FIRST
LIMIT 1;