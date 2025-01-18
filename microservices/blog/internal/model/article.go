package model

import (
	"database/sql"
	"strconv"
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
	Deleted     bool
}

func (a Article) GetStrID() string {
	return strconv.Itoa(a.ID)
}

type ArticleModel struct {
	DB *sql.DB
}

func (m *ArticleModel) GetRange(fromID, count int) ([]Article, error) {
	rows, err := m.DB.Query(
		"SELECT id, title, content, content_html, created_at FROM article where id < $1 AND NOT deleted ORDER BY id DESC LIMIT $2",
		fromID, count,
	)
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
		title, content, m.ContentToHTML(content),
	)
	return err
}

func (m *ArticleModel) Delete(articleID int) error {
	_, err := m.DB.Exec(
		"UPDATE article SET deleted = true WHERE id = $1",
		articleID,
	)
	return err
}

func (m *ArticleModel) ContentToHTML(content string) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.Attributes
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(content))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.LazyLoadImages
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}
