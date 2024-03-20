package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"nurlashko.dev/blog/internal"
	"nurlashko.dev/blog/internal/client"
	"nurlashko.dev/blog/internal/middleware"
	"nurlashko.dev/blog/internal/models"
	"nurlashko.dev/blog/internal/views/article"
	"nurlashko.dev/blog/internal/views/user"
)

func main() {
	config, err := internal.ParseConfig()
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}
	mux := http.NewServeMux()

	am := models.ArticleModel{DB: client.GetDB(config)}
	auth := client.NewAuthClient()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		articles, err := am.All()
		if err != nil {
			log.Printf("error fetching articles: %v", err)
		}
		err = article.ShowIndex(articles).Render(r.Context(), w)
		if err != nil {
			log.Printf("error rendering: %v", err)
		}
	})

	mux.HandleFunc("/article/create", func(w http.ResponseWriter, r *http.Request) {
		toggle := false
		if r.URL.Query().Get("toggle") == "true" {
			toggle = true
		}
		err = article.CreateArticle(toggle).Render(r.Context(), w)
		if err != nil {
			log.Printf("error rendering: %v", err)
		}
	})

	mux.HandleFunc("/article/preview", func(w http.ResponseWriter, r *http.Request) {
		preview := models.Article{
			ID:        42,
			Title:     r.FormValue("title"),
			Content:   r.FormValue("content"),
			CreatedAt: time.Now(),
		}
		err = article.ArticleRow(preview).Render(r.Context(), w)
		if err != nil {
			log.Printf("error rendering: %v", err)
		}
	})

	mux.HandleFunc("/article", middleware.AuthenticationMiddleware(auth, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			title := r.FormValue("title")
			content := r.FormValue("content")
			if len(title) == 0 {
				http.Error(w, "Error: title is too short", http.StatusBadRequest)
				return
			}
			err := am.Insert(title, content)
			if err != nil {
				http.Error(w,
					fmt.Sprintf("failed to insert article: %s", err.Error()), http.StatusBadRequest)
				return
			}
			_ = article.CreateArticle(false).Render(r.Context(), w)
		}
	}))

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			err := user.LoginModal().Render(r.Context(), w)
			if err != nil {
				log.Printf("error rendering: %v", err)
			}
		} else {
			ghToken := r.FormValue("gh_token")
			token, err := auth.GetClientToken(ghToken)
			if err != nil {
				http.Error(w, "failed to get token ", http.StatusUnauthorized)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "x-auth-token",
				Value:    token,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			})
			w.WriteHeader(http.StatusOK)
		}
	})

	fmt.Println("Listening on :8000")
	if err := http.ListenAndServe("0.0.0.0:8000", mux); err != nil {
		log.Printf("error listening: %v", err)
	}
}
