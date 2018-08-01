package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	ID    string `json"id"`
	Fname string `json"fmane"`
	Lname string `json"lname"`
	Age   string `json"Age"`
}

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
}

func ConnectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/patient")
}
