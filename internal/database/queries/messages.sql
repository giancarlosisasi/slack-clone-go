-- name: CreateMessage :one
INSERT INTO messages (user_id, room_id, content, message_type)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, room_id, content, message_type, created_at, updated_at;

-- name: GetRoomMessages :many
-- CRITICAL PERFORMANCE QUERY - Gets recent messages with user info
-- Requires index: CREATE INDEX idx_messages_room_created ON messages(room_id, created_at DESC);
-- this composite index supports both WHERE and ORDER BY clauses
SELECT
  m.id,
  m.user_id,
  m.room_id,
  m.content,
  m.message_type,
  m.created_at,
  u.first_name,
  u.last_name
FROM messages m
JOIN users u on m.user_id = u.id
WHERE m.room_id = $1
ORDER BY m.created_at DESC
LIMIT $2;

-- name: GetMessagesSince :many
-- For real-time syncing - gets messages after a timestamp
-- Uses name composite index as GetRoomMessages
-- Returns in chronological order
SELECT
  m.id,
  m.user_id,
  m.room_id,
  m.content,
  m.message_type,
  m.created_at,
  u.first_name,
  u.last_name
FROM messages m
JOIN users u on m.user_id = u.id
WHERE m.room_id = $1 AND m.created_at > $2
ORDER BY m.created_at DESC;

-- name: GetLatestMessageByRoom :one
-- Gets the most recent message in a room (for room previews)
-- Uses same composite index, just LIMIT 1
SELECT
  m.id,
  m.user_id,
  m.room_id,
  m.content,
  m.message_type,
  m.created_at,
  u.first_name,
  u.last_name
FROM messages m
JOIN users u on m.user_id = u.id
WHERE m.room_id = $1
ORDER BY m.created_at DESC
LIMIT 1;