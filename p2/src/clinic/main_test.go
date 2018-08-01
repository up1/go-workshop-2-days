package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostNewPatient_Success_With_Code_201(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := setupRouter()

	body := bytes.NewBuffer([]byte(`{\"firstname\":"ployploy",\"lastname":\"tobunrueang",\"age":\"23"}`))
	request, err := http.NewRequest("POST", "/api/v1/patient/", body)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Created new patient failed %d", err)
	}
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, request)

	if response.Code != 201 {
		t.Errorf("/api/v1/patient failed %d", response.Code)
	}
}
