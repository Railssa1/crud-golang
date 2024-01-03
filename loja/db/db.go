package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectToBd() *sql.DB {
	connection := "user=postgres dbname=loja password=123456 host=localhost sslmode=disable"
	if db, err := sql.Open("postgres", connection); err != nil {
		panic(err.Error())
	} else {
		return db
	}
}
