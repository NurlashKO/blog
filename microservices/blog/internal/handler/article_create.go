package handler

import (
	"database/sql"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"

	auth "github.com/NurlashKO/blog/microservices/auth/src/client"
	"github.com/NurlashKO/blog/microservices/blog/internal/middleware"
	"github.com/NurlashKO/blog/microservices/blog/internal/model"
	"github.com/NurlashKO/blog/microservices/blog/internal/view/article"
)

func ArticleRangeGET(db *sql.DB) http.HandlerFunc {
	articleModel := model.ArticleModel{DB: db}
	return func(w http.ResponseWriter, r *http.Request) {
		var fromID int
		var err error
		if r.URL.Query().Has("fromID") {
			fromID, err = strconv.Atoi(r.URL.Query().Get("fromID"))
			if err != nil {
				slog.Error("error parsing fromID: %v", err)
				http.Error(w, "error parsing fromID", http.StatusBadRequest)
				return
			}
		} else {
			fromID = math.MaxInt
		}
		articles, err := articleModel.GetRange(fromID, 5)
		if err != nil {
			slog.Error("error fetching articles: %v", err)
		}
		err = article.ArticleList(articles).Render(r.Context(), w)
		if err != nil {
			slog.Error("error rendering: %v", err)
		}
	}
}

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

func ArticleCreatePOST(auth *auth.AuthClient, db *sql.DB) http.HandlerFunc {
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
