package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/feezyhendrix/go-hls-server/internal/db"
	"github.com/feezyhendrix/go-hls-server/pkg/models"
)

func GetAllPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	playlists := []models.Playlist{}
	err := db.DB.Select(&playlists, "SELECT * FROM playlists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(playlists)
}
