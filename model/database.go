package model

import (
	"database/sql"
	"fmt"
)

func openDatabase() *sql.DB {
	dsn := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"

	db, err := sql.Open("postgres", dsn)
	handleError(err)
	return db
}

func closeDatabase() {

}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
