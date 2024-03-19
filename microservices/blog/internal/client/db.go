package client

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	connectionString = "postgres://nurlashko:tmp@database:5432/blog?sslmode=disable"
)

func GetDB() *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}
