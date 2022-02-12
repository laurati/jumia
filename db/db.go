package db

import (
	"database/sql"
	"log"
)

func DbConect() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/sample.db")
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		log.Println("Database connected!")
	}

	return db

}
