package handler

import (
	"github.com/gin-gonic/gin"
)

type Patient struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
}

type PatientResponse struct {
	PatientID string `json:"patientid"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func CreateNewPatient(c *gin.Context) {

}
