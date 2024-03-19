package main

import (
	"fmt"
	"log"
	"net/http"

	"nurlashko.dev/blog/internal"
	"nurlashko.dev/blog/internal/client"
	"nurlashko.dev/blog/internal/models"
	"nurlashko.dev/blog/internal/views/article"
)

func main() {
	config, err := internal.ParseConfig()
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}
	am := models.ArticleModel{DB: client.GetDB(config)}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x, err := am.All()
		if err != nil {
			log.Printf("error fetching articles: %v", err)
		}
		err = article.ShowIndex(x).Render(r.Context(), w)
		if err != nil {
			log.Printf("error rendering: %v", err)
		}
	})

	fmt.Println("Listening on :8000")
	if err := http.ListenAndServe("0.0.0.0:8000", mux); err != nil {
		log.Printf("error listening: %v", err)
	}
}
