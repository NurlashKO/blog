package static

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed html/*
var HtmlTemplates embed.FS

func GetPages() fs.FS {
	f, err := fs.Sub(HtmlTemplates, "html")
	if err != nil {
		log.Fatal(err)
	}
	return f
}
