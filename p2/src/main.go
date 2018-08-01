package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	ID       string
	UserName string
	LastName string
	Age      int
}

func main() {
	http.HandleFunc("/patient", CreatePatient)
	log.Println("Server running on port")
	http.ListenAndServe(":3000", nil)

}

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		username := r.FormValue("patient_username")
		lastname := r.FormValue("patient_lastname")
		age := r.FormValue("patient_age")
		insertForm, err := db.Prepare("INSERT INTO patient(patient_username, patient_lastname,patient_age) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insertForm.Exec(username, lastname, age)
		log.Println("INSERT: Name: " + username + " | Lastname: " + lastname + " | Age: " + age)
	}
	defer db.Close()

	fmt.Fprintf(w, "Create new patient")

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "sckshuhari"
	dbName := "clinic"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
