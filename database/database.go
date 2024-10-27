package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=DracFiendMG dname=users sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}