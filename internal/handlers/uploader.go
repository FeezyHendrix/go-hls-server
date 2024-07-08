package handlers

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/feezyhendrix/go-hls-server/internal/db"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
)

func VideoUploader(w http.ResponseWriter, r *http.Request) {
	// Grab the file
	file, header, err := r.FormFile("file")
	// Do the error not nil dance
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	videoID := ulid.Make().String()
	outputPath := filepath.Join(
		"public",
		"playlists",
		videoID,
	)

	err = exec.Command("mkdir", "-p", outputPath).Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	inputFilePath := filepath.Join(outputPath, header.Filename)
	outFile, err := os.Create(inputFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer outFile.Close()

	_, err = outFile.ReadFrom(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cmd := exec.Command("ffmpeg", "-i", inputFilePath, "-codec: copy", "-start_number", "0", "-hls_time", "10", "-hls_list_size", "0", "-f", "hls", filepath.Join(outputPath, "index.m3u8"))
	err = cmd.Run()
	if err != nil {
		logrus.Info(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.DB.Exec("INSERT INTO playlists (id, video_id, name, path) VALUES ($1, $2, $3, $4)", videoID, header.Filename, header.Filename, outputPath)
	if err != nil {
		logrus.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Upload and processing complete"))
}
