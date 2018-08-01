package member

import (
	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	PatientId string `json:"patient_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}
