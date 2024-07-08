package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func GetSegmentHandler(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "videoId")
	name := chi.URLParam(r, "name")

	segmentPath := filepath.Join("public", "playlists", videoID, name+".ts")
	http.ServeFile(w, r, segmentPath)
}
