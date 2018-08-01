package main

import (
	"github.com/gin-gonic/gin"
)

type patientsResponse struct {
	ID        string `json:"patient_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	router := gin.Default()
	router.POST("/api/v1/patient", routerHandle)
	router.Run(":3000")
}

func routerHandle(c *gin.Context) {
	patient := patientsResponse{
		ID:        "2018-0001",
		Firstname: "นารีนารถ",
		Lastname:  "เนรัญชร",
		Age:       22,
	}
	c.JSON(200, patient)
}
