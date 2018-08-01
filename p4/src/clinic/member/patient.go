package member

import (
	"log"

	"github.com/gin-gonic/gin"
)

type patientsResponse struct {
	ID   string `json:"patient_id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type patientsRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func PatientHandle(c *gin.Context) {
	var patientRequest patientsRequest
	if err := c.ShouldBind(&patientRequest); err == nil {
		c.String(500, "error")
		log.Println("bind error")
	}

	patient := patientsResponse{
		ID:   "2018-0001",
		Name: patientRequest.Firstname + " " + patientRequest.Lastname,
		Age:  patientRequest.Age,
	}
	c.JSON(200, patient)
}
