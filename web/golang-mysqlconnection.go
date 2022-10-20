package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//For localhost: "root:<password>@tcp(127.0.0.1:3306)/<dbname>"
	connectionString := "<username>:<password>@tcp(<HOST>:<port>)/<dbname>"

	//open connection
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//Test Connection
	res, err := db.Query("show tables;")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()

	fmt.Println("Connected Successfully")
}