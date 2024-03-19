package client

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"nurlashko.dev/blog/internal"
)

const (
	connectionStringLocal = "postgres://nurlashko:tmp@localhost:5432/blog?sslmode=disable"
	connectionStringProd  = "postgres://nurlashko:tmp@database:5432/blog?sslmode=disable"
)

func GetDB(config internal.Config) *sql.DB {
	connectionString := connectionStringProd
	if config.Debug {
		connectionString = connectionStringLocal
	}
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}
