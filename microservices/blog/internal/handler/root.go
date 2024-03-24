package handler

import (
	"database/sql"
	"log/slog"
	"net/http"

	"nurlashko.dev/blog/internal/model"
	"nurlashko.dev/blog/internal/view/article"
)

func RootHandler(db *sql.DB) http.HandlerFunc {
	articleModel := model.ArticleModel{DB: db}
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := articleModel.All()
		if err != nil {
			slog.Info("error fetching articles: %v", err)
		}
		err = article.ShowIndex(articles).Render(r.Context(), w)
		if err != nil {
			slog.Info("error rendering: %v", err)
		}
	}
}
