package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("api/v1/patient", routerHandler)
	router.Run(":8080")
}

type patientResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type patientRequest struct {
	ID    string `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}

func routerHandler(c *gin.Context) {
	var patientRequest patientRequest
	err := c.ShouldBind(&patientRequest)
	if err != nil {
		log.Println("bind error")
	}
	patient := patientResponse{
		Id:   "2018-0001",
		Name: patientRequest.Fname + patientRequest.Lname,
		Age:  patientRequest.Age,
	}
	c.JSON(201, patient)
}
