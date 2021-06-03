package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:00@tcp(127.0.0.1:3306)/login")
	if err != nil {
		panic(err.Error())
	}
	return db
}
