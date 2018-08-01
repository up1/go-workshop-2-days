package handler

import (
	"github.com/gin-gonic/gin"
)

// Request รูปแบบของสิ่งที่จะรับเข้ามา
type Request struct {
	FirstName string
	LastName  string
	Age       int
}

// Response ของที่จะส่งกลับไปที่ user
type Response struct {
	PatientID string
	FirstName string
	LastName  string
	Age       int
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
	var patient Patient
	c.BindJSON(&patient)
	c.JSON(200, patient)
}
