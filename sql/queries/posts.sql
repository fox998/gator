
-- name: CreatePost :one
INSERT INTO posts (created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetPostsForUser :many
WITH user_feeds AS (
    SELECT feed_follows.feed_id AS id
    FROM feed_follows
    WHERE feed_follows.user_id = (SELECT id FROM users WHERE name = $1)
)
SELECT posts.*
FROM posts
WHERE posts.feed_id IN (SELECT id FROM user_feeds)
ORDER BY posts.published_at DESC
LIMIT $2;
