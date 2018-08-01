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

type Api struct {
	PatientService PatientService
}

type PatientService struct {
	DBSession *mgo.Session
}

func main() {
	DBSession, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("Cannot connect database ", err.Error())
		return
	}
	patientService := PatientService{DBSession: DBSession}
	api := Api{PatientService: patientService}
	http.HandleFunc("/v1/patients", api.CreatePatientHandler)

	http.ListenAndServe(":3000", nil)

}

func (api Api) CreatePatientHandler(writer http.ResponseWriter, request *http.Request) {
	var patient Patient
	err := json.NewDecoder(request.Body).Decode(&patient)
	if err != nil {
		http.Error(writer, "Not Found", http.StatusNotFound)
		return
	}
	newPatient, err := api.PatientService.InsertPatient(patient)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(newPatient)
}

func (service PatientService) InsertPatient(patient Patient) (Patient, error) {
	var newPatient Patient
	patient.PatientID, _ = service.GeneratePatientID()
	err := service.DBSession.DB("hospital_somkiat").C("patients").Insert(patient)
	if err != nil {
		return patient, err
	}
	service.UpdatePatientCount()
	err = service.DBSession.DB("hospital_somkiat").C("patients").Find(bson.M{"patientID": patient.PatientID}).One(&newPatient)
	return newPatient, err
}

func (service PatientService) GeneratePatientID() (string, error) {
	var patientCount PatientCount
	err := service.DBSession.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	if err != nil && err.Error() == "not found" {
		patientCount.Count = 0
		service.DBSession.DB("hospital_somkiat").C("patientsCount").Insert(&patientCount)
	}
	patientCount.Count++
	return service.FormatPatientID(time.Now().Year(), patientCount.Count), nil
}

func (service PatientService) UpdatePatientCount() {
	var patientCount PatientCount
	service.DBSession.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	service.DBSession.DB("hospital_somkiat").C("patientsCount").UpdateId(patientCount.ID, bson.M{"$inc": bson.M{"count": 1}})
}

func (service PatientService) FormatPatientID(year, number int) string {
	return fmt.Sprintf("%d-%04d", year, number)
}
