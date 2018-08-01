package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_main_CreateNewPatient_Success_With_201(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()

	body := bytes.NewBuffer([]byte(`{"name":"Mananchaya","lastname":"Amornpalang","age":"24"}`))
	req, err := http.NewRequest("POST", "/api/v1/patients", body)
	req.Header.Set("Content-Type", "application.json")
	if err != nil {
		t.Errorf("Create new patient failed with error %d", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 201 {
		t.Errorf("api/v1/patient failed with error %d", resp.Code)
	}

}
