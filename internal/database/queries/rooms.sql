-- name: CreateRoom :one
INSERT INTO rooms (name)
VALUES ($1)
RETURNING id, name, created_at;

-- name: GetRoomByName :one
SELECT id, name, created_at
FROM rooms
WHERE name = $1;

-- name: GetRoomByID :one
SELECT id, name, created_at
FROM rooms
WHERE id = $1;

-- name: GetAllRooms :many
SELECT id, name, created_at
FROM rooms
ORDER BY name ASC;

-- name: AddUserToRoom :exec
-- Adds user to room, ignores if already exists (prevents duplicated)
INSERT INTO room_members (room_id, user_id)
VALUES ($1, $2)
ON CONFLICT (room_id, user_id) DO NOTHING;

-- name: RemoveUserFromRoom :exec
DELETE FROM room_members
WHERE room_id = $1 AND user_id = $2;

-- name: GetRoomMembers :many
-- Gets all users in a room with join timestamp
-- Uses foreign key indexes automatically
SELECT sqlc.embed(u), rm.joined_at
FROM room_members rm
JOIN users u ON rm.user_id = u.id
WHERE rm.room_id = $1
ORDER BY rm.joined_at ASC;