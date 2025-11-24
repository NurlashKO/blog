package handler

import (
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"github.com/NurlashKO/blog/microservices/blog/internal/model"
	"github.com/NurlashKO/blog/microservices/blog/internal/view/article"
)

func ArticlePreview(db *sql.DB) http.HandlerFunc {
	articleModel := model.ArticleModel{DB: db}
	return func(w http.ResponseWriter, r *http.Request) {
		preview := model.Article{
			ID:          42,
			Title:       r.FormValue("title"),
			Content:     r.FormValue("content"),
			ContentHtml: articleModel.ContentToHTML(r.FormValue("content")),
			CreatedAt:   time.Now(),
		}
		err := article.ArticleRow(preview, true).Render(r.Context(), w)
		if err != nil {
			slog.Info("error rendering: %v", err)
		}
	}
}
