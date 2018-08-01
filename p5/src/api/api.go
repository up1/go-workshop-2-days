package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Patient struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

type PatientResponse struct {
	PatientID string `json:"patientid"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func PostInfomationPatientID(respons http.ResponseWriter, request *http.Request) {
	patient := Patient{}
	err := json.NewDecoder(request.Body).Decode(&patient)

	if err != nil {
		http.Error(respons, err.Error(), http.StatusInternalServerError)
		return
	}

	PatientResponse := patient

	newPatientResponseJson, _ := json.Marshal(PatientResponse)
	respons.Write(newPatientResponseJson)

}
func ConnectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/patient")
}
