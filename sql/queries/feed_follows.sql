-- name: CreateFeedFollow :one
WITH 
feed AS (SELECT * FROM feeds WHERE url = $1),
follower AS (SELECT * FROM users WHERE name = $2),
inserted AS (
    INSERT INTO feed_follows (created_at, updated_at, feed_id, user_id)
    VALUES (NOW(), NOW(), (SELECT id FROM feed), (SELECT id FROM follower))
    RETURNING *
)
SELECT feed.title, feed.url, follower.name AS follower_name
FROM inserted 
JOIN feed ON inserted.feed_id = feed.id
JOIN follower ON inserted.user_id = follower.id;

-- name: CreateFeedFollowIds :one
INSERT INTO feed_follows (created_at, updated_at, feed_id, user_id)
VALUES (NOW(), NOW(), $1, $2)
RETURNING *;

-- name: GetFeedFollowsForUser :many
WITH followed AS (
    SELECT feed_follows.feed_id AS id
    FROM feed_follows
    WHERE feed_follows.user_id = (SELECT id FROM users WHERE name = $1)
)
SELECT feeds.title, feeds.url
FROM feeds
WHERE feeds.id IN (SELECT id FROM followed);

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE user_id = (SELECT id FROM users WHERE name = $1) AND feed_id = (SELECT id FROM feeds WHERE url = $2);
