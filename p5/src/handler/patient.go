package handler

import (
	"github.com/gin-gonic/gin"
)

//PatientRequest keep infomation
type PatientRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      string `json:"age"`
}

//PatientResponse response infomation
type PatientResponse struct {
	PatientID string `json:"patientID"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Age       string `json:"age"`
}

// CreateNewPatient response Patient infomations
func CreateNewPatient(c *gin.Context) {

	var patientRequest PatientRequest
	c.ShouldBind(&patientRequest)
	patientResp := PatientResponse{
		PatientID: "2018-0001",
		Name:      patientRequest.Name,
		Lastname:  patientRequest.Lastname,
		Age:       patientRequest.Age,
	}
	c.JSON(201, patientResp)

}
