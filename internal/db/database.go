package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
)

func CreateDatabase() (err error) {
	DBConn, err = sql.Open("mysql", "user:password@/db")
	defer DBConn.Close()
	if err != nil {
		return err
	}

	return err
}