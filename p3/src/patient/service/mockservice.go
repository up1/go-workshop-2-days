package service

import (
	"patient/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MockPatientService struct {
}

func (mps MockPatientService) InsertPatient(patient model.Patient) (model.Patient, error) {
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Aug-1")
	patient.ID = bson.NewObjectIdWithTime(fixedTime)
	patient.PatientID = "2018-0001"
	return patient, nil
}
