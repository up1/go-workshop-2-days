package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

// NewPatientRequest รูปแบบของสิ่งที่จะรับเข้ามา
type NewPatientRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

// NewPatientResponse ของที่จะส่งกลับไปที่ user
type NewPatientResponse struct {
	PatientID string `json:"patient_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

// Patient struct ของผู้ป่วย
type Patient struct {
	ID        string `json:"id"`
	PatientID string `json:"patient_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

// CreatePatient สร้างผู้ป่วย
func CreatePatient(c *gin.Context) {
	patientRequest := new(NewPatientRequest)
	err := c.ShouldBind(&patientRequest)
	if err != nil {
		log.Println("error")
	}
	patient := NewPatientResponse{
		PatientID: "2018-0001",
		FirstName: patientRequest.FirstName,
		LastName:  patientRequest.LastName,
		Age:       patientRequest.Age,
	}
	c.JSON(201, patient)

}
