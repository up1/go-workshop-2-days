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
		var patient Patient
		err := json.NewDecoder(request.Body).Decode(&patient)
		if err != nil {
			http.Error(writer, "Not Found", http.StatusNotFound)
			return
		}
		newPatient, err := InsertPatient(DBSession, patient)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode(newPatient)
	})

	http.ListenAndServe(":3000", nil)

}

func InsertPatient(session *mgo.Session, patient Patient) (Patient, error) {
	var newPatient Patient
	patient.PatientID, _ = GeneratePatientID(session)
	err := session.DB("hospital_somkiat").C("patients").Insert(patient)
	if err != nil {
		return patient, err
	}
	UpdatePatientCount(session)
	err = session.DB("hospital_somkiat").C("patients").Find(bson.M{"patientID": patient.PatientID}).One(&newPatient)
	return newPatient, err
}

func GeneratePatientID(session *mgo.Session) (string, error) {
	var patientCount PatientCount
	err := session.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	if err != nil && err.Error() == "not found" {
		patientCount.Count = 0
		session.DB("hospital_somkiat").C("patientsCount").Insert(&patientCount)
	}
	patientCount.Count++
	return FormatPatientID(time.Now().Year(), patientCount.Count), nil
}

func UpdatePatientCount(session *mgo.Session) {
	var patientCount PatientCount
	session.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	session.DB("hospital_somkiat").C("patientsCount").UpdateId(patientCount.ID, bson.M{"$inc": bson.M{"count": 1}})
}

func FormatPatientID(year, number int) string {
	return fmt.Sprintf("%d-%04d", year, number)
}
