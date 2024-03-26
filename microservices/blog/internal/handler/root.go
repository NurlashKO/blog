package handler

import (
	"database/sql"
	"log/slog"
	"math"
	"net/http"
	"slices"
	"strconv"

	"nurlashko.dev/blog/internal/model"
	"nurlashko.dev/blog/internal/view/article"
)

func RootHandler(db *sql.DB) http.HandlerFunc {
	articleModel := model.ArticleModel{DB: db}
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := articleModel.GetRange(math.MaxInt32, 5)
		if err != nil {
			slog.Error("error fetching articles: %v", err)
		}
		// Add specific article to the top if requested by GET parameter
		if r.URL.Query().Has("articleID") {
			articleID, err := strconv.Atoi(r.URL.Query().Get("articleID"))
			if err == nil {
				a, err := articleModel.GetRange(articleID+1, 1)
				if err == nil {
					if len(a) != 1 || a[0].ID != articleID {
						slog.Error("article not found: %v", articleID)
					} else {
						articles = slices.DeleteFunc(articles, func(i model.Article) bool { return i.ID == a[0].ID })
						articles = slices.Concat(a, articles)
					}
				} else {
					slog.Error("error fetching article: %v", err)
				}
			} else {
				slog.Error("error parsing articleID: %v", err)
			}
		}
		err = article.ShowIndex(articles).Render(r.Context(), w)
		if err != nil {
			slog.Info("error rendering: %v", err)
		}
	}
}
