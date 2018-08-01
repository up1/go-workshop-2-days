package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"patient/model"
	"patient/router"
	"testing"

	"github.com/labstack/echo"
)

func Test_CreatePatient_Should_Be_NewPatient(t *testing.T) {
	var actualPatient model.NewPatientResponse
	patient := model.NewPatientRequest{
		FirstName: "Apipol",
		LastName:  "Sukgler",
		Age:       25,
	}
	patientJson, _ := json.Marshal(patient)
	expectedPatient := model.NewPatientResponse{
		PatientID: "2018-0001",
		FirstName: "Apipol",
		LastName:  "Sukgler",
		Age:       25,
	}

	request := httptest.NewRequest("POST", "/api/v1/patients", bytes.NewBuffer(patientJson))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	route := router.SetupRouteEcho()
	route.ServeHTTP(recorder, request)
	response := recorder.Result()

	json.NewDecoder(response.Body).Decode(&actualPatient)
	if response.StatusCode != 201 {
		t.Errorf("expected %d but it got %d", 201, response.StatusCode)
	}
	if actualPatient != expectedPatient {
		t.Errorf("expected %v but it got %v", expectedPatient, actualPatient)
	}
}
