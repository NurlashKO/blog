package model

import (
	"database/sql"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Article struct {
	ID          int
	Title       string
	Content     string
	ContentHtml string
	CreatedAt   time.Time
}

type ArticleModel struct {
	DB *sql.DB
}

func (m *ArticleModel) All() ([]Article, error) {
	rows, err := m.DB.Query("SELECT id, title, content, content_html, created_at FROM article ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []Article
	for rows.Next() {
		article := Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.ContentHtml, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (m *ArticleModel) Insert(title, content string) error {
	_, err := m.DB.Exec(
		"INSERT INTO article (title, content, content_html) VALUES ($1, $2, $3)",
		title, content, m.ContentToHTML(content))
	return err
}

func (m *ArticleModel) ContentToHTML(content string) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(content))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}
