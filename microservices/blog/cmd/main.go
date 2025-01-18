package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

	auth "nurlashko.dev/auth/client"
	"nurlashko.dev/blog/internal"
	"nurlashko.dev/blog/internal/client"
	"nurlashko.dev/blog/internal/handler"
)

type BlogApp struct {
	auth   *auth.AuthClient
	config internal.Config
	db     *sql.DB
}

func NewBlogApp() *BlogApp {
	config, err := internal.ParseConfig()
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return &BlogApp{
		auth:   auth.NewAuthClient(config.Debug),
		db:     client.GetDB(config),
		config: config,
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	mux := http.NewServeMux()
	app := NewBlogApp()

	mux.HandleFunc("GET /", handler.RootHandler(app.db))

	mux.HandleFunc("GET /article", handler.ArticleRangeGET(app.db))
	mux.HandleFunc("GET /article/create", handler.ArticleCreateGET())
	mux.HandleFunc("POST /article/create", handler.ArticleCreatePOST(app.auth, app.db))
	mux.HandleFunc("DELETE /article/delete", handler.ArticleDelete(app.auth, app.db))

	mux.HandleFunc("PUT /article/preview", handler.ArticlePreview(app.db))

	mux.HandleFunc("GET /login", handler.LoginGET())

	mux.HandleFunc("POST /login", handler.LoginPOST(app.auth, app.config.Debug))

	slog.Info("Listening on :8000")
	if err := http.ListenAndServe("0.0.0.0:8000", mux); err != nil {
		slog.Error("error listening: %v", err)
	}
}
