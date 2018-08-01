package member

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	Patient_id string `json:"patient_id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Age        int    `json:"age"`
}

func Member(patient_id, firstname, lastname string, age int) string {
	member := ConnectDB()

	return fmt.Sprintf("%s %s", member, patient_id, firstname, lastname, age)
}

func GetPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msg := Patient{
		Patient_id: "1995-0001",
		Firstname:  "นารีนารถ",
		Lastname:   "เนรัญชร",
		Age:        22,
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}

func ConnectDB() (db *sql.DB) {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/clinic")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	if r.Method == "POST" {
		patient_id := r.FormValue("1995-0002")
		firstname := r.FormValue("ภาณุมาศ")
		lastname := r.FormValue("แสนโท")
		age := 22
		insForm, err := db.Prepare("INSERT INTO patient(patient_id,firstname,lastname, age) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(patient_id, firstname, lastname, age)
		log.Println("INSERT: Patient_id: " + patient_id + " | name: " + firstname + lastname)
	}
	defer db.Close()
}
