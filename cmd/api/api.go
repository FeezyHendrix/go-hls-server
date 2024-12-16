package api

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/feezyhendrix/go-hls-server/internal/db"
	"github.com/feezyhendrix/go-hls-server/internal/handlers"
	"github.com/feezyhendrix/go-hls-server/internal/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Run() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	err = db.InitDB(dsn)
	if err != nil {
		logrus.Fatal(err)
	}

	r := router()
	logrus.Info("Running on port " + port)
	http.ListenAndServe(port, r)
}

func router() *chi.Mux {
	r := chi.NewRouter()

	logrusLogger := logrus.New()
	logrusLogger.Formatter = &logrus.JSONFormatter{}

	r.Use(middleware.RequestID)
	r.Use(logRequest)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	/// Templates
	r.Get("/", views.IndexView)

	r.Post("/api/v1/upload", handlers.VideoUploader)
	r.Get("/api/v1/playlists", handlers.GetAllPlaylistsHandler)
	r.Get("/api/v1/segment/{videoId}/{name}.ts", handlers.GetSegmentHandler)

	return r
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entry := logrus.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"remote": r.RemoteAddr,
		})
		entry.Info("request started")
		start := time.Now()

		next.ServeHTTP(w, r)

		entry.WithField("duration", time.Since(start)).Info("request completed")
	})
}
