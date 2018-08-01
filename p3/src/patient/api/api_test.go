package api_test

import (
	"encoding/json"
	"net/http/httptest"
	. "patient/api"
	"patient/model"
	"patient/service"
	"strings"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func Test_CreatePatientHandler_Input_Firstname_Apipol_Lastname_Sukgler_Age_25_Should_Be_PatientID_2018_0001_With_Sample_informations(t *testing.T) {
	var actualPatient model.Patient
	patient := model.Patient{
		FirstName: "Apipol",
		LastName:  "Sukgler",
		Age:       25,
	}
	fixedTime, _ := time.Parse("2006-Jan-02", "2018-Aug-1")
	expectedPatient := model.Patient{
		ID:        bson.NewObjectIdWithTime(fixedTime),
		PatientID: "2018-0001",
		FirstName: "Apipol",
		LastName:  "Sukgler",
		Age:       25,
	}
	patientJson, _ := json.Marshal(patient)

	api := Api{
		PatientService: &service.MockPatientService{},
	}
	request := httptest.NewRequest("POST", "/v1/patients", strings.NewReader(string(patientJson)))
	recorder := httptest.NewRecorder()
	api.CreatePatientHandler(recorder, request)
	response := recorder.Result()
	json.NewDecoder(response.Body).Decode(&actualPatient)
	if actualPatient != expectedPatient {
		t.Errorf("expected %v but it got %v", expectedPatient, actualPatient)
	}

}
