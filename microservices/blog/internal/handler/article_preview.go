package handler

import (
	"log/slog"
	"net/http"
	"time"

	"nurlashko.dev/blog/internal/model"
	"nurlashko.dev/blog/internal/view/article"
)

func ArticlePreview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		preview := model.Article{
			ID:        42,
			Title:     r.FormValue("title"),
			Content:   r.FormValue("content"),
			CreatedAt: time.Now(),
		}
		err := article.ArticleRow(preview, true).Render(r.Context(), w)
		if err != nil {
			slog.Info("error rendering: %v", err)
		}
	}
}
