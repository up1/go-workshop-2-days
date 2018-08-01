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

func ConnectDB() string {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/clinic")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var member string
	err = db.QueryRow("SELECT patient_id FROM patient WHERE patient_id = ?", 1995-0001).Scan(&member)
	if err != nil {
		log.Fatal(err)
	}
	return member
}
