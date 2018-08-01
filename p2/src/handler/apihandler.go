package handler

import (
	"fmt"
	"net/http"

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
	ID        string `json: "id"`
	PatientID string `json: "patient_id"`
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
	Age       int    `json: "age"`
}

func CreatePatient(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	fmt.Printf("Create new patient")

}
