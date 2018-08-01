package main

import (
	"clinic/router"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_RouterCreatePatient_Should_Be_Status_Code_201(t *testing.T) {
	patient := router.PatientResponse{
		PatientID: "2018-0001",
		Name:      "panumars seanto",
		Age:       25,
	}
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()
	body, _ := json.Marshal(patient)
	strinBody := strings.NewReader(string(body))
	request, err := http.NewRequest("POST", "/api/v1/patient", strinBody)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Create new patient failed with error %d.", err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)
	if response.Code != 201 {
		t.Errorf("expect code is 201 but got %d", response.Code)
	}
}
