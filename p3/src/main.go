package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Patient struct {
	ID        string `json:"id,omitempty"`
	PatientID string `json:"patientID,omitempty"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/v1/patients", func(writer http.ResponseWriter, request *http.Request) {
		var patient Patient
		err := json.NewDecoder(request.Body).Decode(&patient)
		if err != nil {
			http.Error(writer, "Not Found", http.StatusNotFound)
		}

		json.NewEncoder(writer).Encode(patient)
	})
}

func GeneratePatientID() string {
	return strconv.Itoa(time.Now().Year())
}
