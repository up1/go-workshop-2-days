package handler

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	FirstName string
	LastName  string
	Age       int
}

type Response struct {
	PatientID string
	FirstName string
	LastName  string
	Age       int
}

type Patient struct {
	ID        string `json:"id"`
	PatientID string `json:"patient_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func CreatePatient(c *gin.Context) {
	var patient Patient
	c.BindJSON(&patient)
	c.JSON(200, patient)
}
