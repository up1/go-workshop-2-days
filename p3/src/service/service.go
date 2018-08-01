package service

import (
	"fmt"
	"model"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PatientService interface {
	InsertPatient(patient model.Patient) (model.Patient, error)
}
type MongoPatientService struct {
	DBSession *mgo.Session
}

func NewMongoPatientService(session *mgo.Session) MongoPatientService {
	return MongoPatientService{
		DBSession: session,
	}
}
func (service MongoPatientService) InsertPatient(patient model.Patient) (model.Patient, error) {
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

func (service MongoPatientService) GeneratePatientID() (string, error) {
	var patientCount model.PatientCount
	err := service.DBSession.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	if err != nil && err.Error() == "not found" {
		patientCount.Count = 0
		service.DBSession.DB("hospital_somkiat").C("patientsCount").Insert(&patientCount)
	}
	patientCount.Count++
	return FormatPatientID(time.Now().Year(), patientCount.Count), nil
}

func (service MongoPatientService) UpdatePatientCount() {
	var patientCount model.PatientCount
	service.DBSession.DB("hospital_somkiat").C("patientsCount").Find(nil).One(&patientCount)
	service.DBSession.DB("hospital_somkiat").C("patientsCount").UpdateId(patientCount.ID, bson.M{"$inc": bson.M{"count": 1}})
}

func FormatPatientID(year, number int) string {
	return fmt.Sprintf("%d-%04d", year, number)
}
