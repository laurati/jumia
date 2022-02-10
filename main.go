package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"

	_ "github.com/mattn/go-sqlite3"
)

type Sample struct {
	Id       int
	Name     string
	Telefone string
}

type Jumia_table struct {
	Country      string
	State        string
	Country_code string
	Phone_num    string
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

	//fmt.Println(sampleArr)

	re := regexp.MustCompile("\\(212\\)?[5-9]\\d{8}$")
	fmt.Println(re.MatchString("(212)520123453"))

	client := Jumia_table{}
	var clientArr []Jumia_table
	var state bool

	for _, v := range sampleArr {
		s := v.Telefone[1:4]
		switch s {
		case "237":
			client.Country = "Cameroon"
			client.Country_code = fmt.Sprintf("+%s", s)
			ra := regexp.MustCompile("\\(237\\) ?[2368]\\d{7,8}$")
			client.Phone_num = v.Telefone
			state = ra.MatchString(v.Telefone)
			client.State = whatIs(state)
		case "251":
			client.Country = "Ethiopia"
			client.Country_code = fmt.Sprintf("+%s", s)
			ra := regexp.MustCompile("\\(251\\) ?[1-59]\\d{8}$")
			client.Phone_num = v.Telefone
			state = ra.MatchString(v.Telefone)
			client.State = whatIs(state)
		case "212":
			client.Country = "Morocco"
			client.Country_code = fmt.Sprintf("+%s", s)
			ra := regexp.MustCompile("\\(212\\) ?[5-9]\\d{8}$")
			client.Phone_num = v.Telefone
			state = ra.MatchString(v.Telefone)
			client.State = whatIs(state)
		case "258":
			client.Country = "Mozambique"
			client.Country_code = fmt.Sprintf("+%s", s)
			ra := regexp.MustCompile("\\(258\\) ?[28]\\d{7,8}$")
			client.Phone_num = v.Telefone
			state = ra.MatchString(v.Telefone)
			client.State = whatIs(state)
		case "256":
			client.Country = "Uganda"
			client.Country_code = fmt.Sprintf("+%s", s)
			ra := regexp.MustCompile("\\(256\\) ?\\d{9}$")
			client.Phone_num = v.Telefone
			state = ra.MatchString(v.Telefone)
			client.State = whatIs(state)
		}

		client = Jumia_table{client.Country, client.State, client.Country_code, client.Phone_num}
		clientArr = append(clientArr, client)
	}

	fmt.Println(clientArr)

}

func whatIs(b bool) string {
	if b == true {
		return "OK"
	}
	return "NOK"
}
