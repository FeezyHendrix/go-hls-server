package models

type Playlist struct {
	ID      string `db: "id" json: "id"`
	VideoID string `db: "video_id" json: "videoId"`
	Name    string `db: "name" json: "name"`
	Path    string `db: "path" json: "path"`
}
