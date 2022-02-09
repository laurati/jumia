package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sample struct {
	Id       int
	Name     string
	Telefone string
}

func main() {

	db, err := sql.Open("sqlite3", "./db/sample.db")
	if err != nil {
		log.Println("Error:", err.Error())
	}

	// query
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		log.Println("Error:", err.Error())
	}

	customer := Sample{}
	var sampleArr []Sample

	for rows.Next() {
		err = rows.Scan(&customer.Id, &customer.Name, &customer.Telefone)
		if err != nil {
			log.Println("Error:", err.Error())
		}

		sampleArr = append(sampleArr, customer)

	}

	rows.Close()

	fmt.Println(sampleArr)

}
