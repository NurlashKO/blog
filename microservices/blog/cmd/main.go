package main

import (
	"fmt"
	"log"
	"net/http"

	"nurlashko.dev/blog/internal/client"
	"nurlashko.dev/blog/internal/models"
	"nurlashko.dev/blog/internal/views/article"
)

func main() {
	am := models.ArticleModel{DB: client.GetDB()}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x, _ := am.All()
		_ = article.ShowIndex(x).Render(r.Context(), w)
	})

	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe("localhost:3000", mux); err != nil {
		log.Printf("error listening: %v", err)
	}
}
