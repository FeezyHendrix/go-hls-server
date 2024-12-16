-- migrations/000001_create_playlists_table.up.sql
CREATE TABLE playlists (
    id TEXT PRIMARY KEY,
    video_id TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    path TEXT NOT NULL
);
