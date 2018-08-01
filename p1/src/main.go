package main

import (
	. "api"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("connect success")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	apiConnect := API{ConnectDB: db}
	http.HandleFunc("/patient/add", apiConnect.AddMember)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ConnectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/patient")
}
