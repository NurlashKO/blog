package models

import (
	"database/sql"
	"time"
)

type Article struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

type ArticleModel struct {
	DB *sql.DB
}

func (m *ArticleModel) All() ([]Article, error) {
	rows, err := m.DB.Query("SELECT * FROM article ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []Article
	for rows.Next() {
		article := Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (m *ArticleModel) Insert(title, content string) error {
	_, err := m.DB.Exec("INSERT INTO article (title, content) VALUES ($1, $2)", title, content)
	return err
}
