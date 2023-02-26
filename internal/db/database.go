package db

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
)

func CreateDatabase() (err error) {
	DBConn, err = sql.Open("mysql", "user:password@tcp(gym-app-db)/db?parseTime=true")
	// defer DBConn.Close()
	if err != nil {
		log.Fatal(err)
	}

	return err
}