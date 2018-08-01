package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const url = "mongodb://localhost:27017"

type PatientCount struct {
	ID    bson.ObjectId `json:"id,omitempty"bson:"_id,omitempty"`
	Count int           `json:"count"bson:"count"`
}

type Patient struct {
	ID        bson.ObjectId `json:"id,omitempty"bson:"_id,omitempty"`
	PatientID string        `json:"patientID,omitempty"bson:"patientID"`
	FirstName string        `json:"firstname"bson:"firstname"`
	LastName  string        `json:"lastname"bson:"lastname"`
	Age       int           `json:"age"bson:"age"`
}

func main() {
	DBSession, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("Cannot connect database ", err.Error())
		return
	}

	http.HandleFunc("/v1/patients", func(writer http.ResponseWriter, request *http.Request) {
		var patient, newPatient Patient
		err := json.NewDecoder(request.Body).Decode(&patient)
		fmt.Println(request.Body)
		if err != nil {
			http.Error(writer, "Not Found", http.StatusNotFound)
			return
		}
		patient.PatientID, _ = GeneratePatientID(DBSession)
		err = DBSession.DB("hospital_somkiat").C("patients").Insert(patient)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		err = DBSession.DB("hospital_somkiat").C("patients").Find(bson.M{"patientID": patient.PatientID}).One(&newPatient)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(writer).Encode(newPatient)
	})

	http.ListenAndServe(":3000", nil)

}

func GeneratePatientID(session *mgo.Session) (string, error) {
	var patientCount PatientCount
	err := session.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	if err != nil && err.Error() == "not found" {
		patientCount.Count = 0
	}
	patientCount.Count++
	return FormatPatientID(time.Now().Year(), patientCount.Count), nil
}

func FormatPatientID(year, number int) string {
	fmt.Println(year, number)
	return fmt.Sprintf("%d-%4d", year, number)
}
