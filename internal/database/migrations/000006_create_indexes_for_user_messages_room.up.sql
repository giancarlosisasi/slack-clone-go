-- Essential performance indexes
-- Email and name are unique so maybe we don't need to create indexes for them?
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE UNIQUE INDEX IF NOT EXISTS idx_rooms_name ON rooms(name);

-- THE most critical index for chat performance
CREATE INDEX IF NOT EXISTS idx_messages_room_created ON messages(room_id, created_at DESC);

-- Foreign key performance indexes
CREATE INDEX IF NOT EXISTS idx_messages_user_id ON messages(user_id);
CREATE INDEX IF NOT EXISTS idx_room_members_user_id ON room_members(user_id);