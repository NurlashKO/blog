package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	auth "nurlashko.dev/auth/client"
	"nurlashko.dev/blog/internal/middleware"
	"nurlashko.dev/blog/internal/model"
	"nurlashko.dev/blog/internal/view/article"
)

func ArticleDelete(auth *auth.AuthClient, db *sql.DB) http.HandlerFunc {
	articleModel := model.ArticleModel{DB: db}
	return middleware.AuthenticationMiddleware(auth, func(w http.ResponseWriter, r *http.Request) {
		articleID, err := strconv.Atoi(r.URL.Query().Get("articleID"))
		if err != nil {
			http.Error(w, "error parsing articleID", http.StatusBadRequest)
			return
		}
		err = articleModel.Delete(articleID)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("failed to delete article: %s", err.Error()), http.StatusBadRequest)
			return
		}
		_ = article.DeleteArticle().Render(r.Context(), w)
	})
}
