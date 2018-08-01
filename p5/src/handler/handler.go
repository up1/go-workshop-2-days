package handler

import (
	"github.com/gin-gonic/gin"
)

type Patient struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      string `json:"age"`
}

type PatientResponse struct {
	PatientID string `json:"patientid"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Age       string `json:"age"`
}

func CreateNewPatient(c *gin.Context) {

	var patient Patient

	//var patientResp PatientResponse

	c.ShouldBind(&patient)
	c.JSON(200, gin.H{
		"patientID": "2018-0001",
		"name":      patient.Name,
		"lastname":  patient.Lastname,
		"age":       patient.Age,
	})

}
