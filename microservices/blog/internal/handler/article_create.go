package handler

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"nurlashko.dev/blog/internal/client"
	"nurlashko.dev/blog/internal/middleware"
	"nurlashko.dev/blog/internal/model"
	"nurlashko.dev/blog/internal/view/article"
)

func ArticleCreateGET() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		toggle := false
		if r.URL.Query().Get("toggle") == "true" {
			toggle = true
		}
		err := article.CreateArticle(toggle).Render(r.Context(), w)
		if err != nil {
			slog.Info("error rendering: %v", err)
		}
	}
}

func ArticleCreatePOST(auth *client.AuthClient, db *sql.DB) http.HandlerFunc {
	articleModel := model.ArticleModel{DB: db}
	return middleware.AuthenticationMiddleware(auth, func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		content := r.FormValue("content")
		if len(title) == 0 {
			http.Error(w, "Error: title is too short", http.StatusBadRequest)
			return
		}
		err := articleModel.Insert(title, content)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("failed to insert article: %s", err.Error()), http.StatusBadRequest)
			return
		}
		_ = article.CreateArticle(false).Render(r.Context(), w)
	})
}
