package service

import (
	"fmt"
	"model"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type IPatientService interface {
	InsertPatient(patient model.Patient) (model.Patient, error)
}
type PatientService struct {
	DBSession *mgo.Session
}

func (service PatientService) InsertPatient(patient model.Patient) (model.Patient, error) {
	var newPatient model.Patient
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
	var patientCount model.PatientCount
	err := service.DBSession.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	if err != nil && err.Error() == "not found" {
		patientCount.Count = 0
		service.DBSession.DB("hospital_somkiat").C("patientsCount").Insert(&patientCount)
	}
	patientCount.Count++
	return service.FormatPatientID(time.Now().Year(), patientCount.Count), nil
}

func (service PatientService) UpdatePatientCount() {
	var patientCount model.PatientCount
	service.DBSession.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	service.DBSession.DB("hospital_somkiat").C("patientsCount").UpdateId(patientCount.ID, bson.M{"$inc": bson.M{"count": 1}})
}

func (service PatientService) FormatPatientID(year, number int) string {
	return fmt.Sprintf("%d-%04d", year, number)
}
