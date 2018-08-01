package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PatientResponse struct {
	PatientID string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
}
type patientRequest struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}

func RouterCreatePatient(c *gin.Context) {
	var patientRequest patientRequest
	err := c.ShouldBind(&patientRequest)
	if err != nil {
		c.String(500, "error")
		log.Println("bind error")
	}
	patient := PatientResponse{
		PatientID: "2018-0001",
		Name:      patientRequest.Fname + " " + patientRequest.Lname,
		Age:       patientRequest.Age,
	}
	c.JSON(201, patient)
}
