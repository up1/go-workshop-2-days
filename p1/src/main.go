package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("api/v1/patient", routerHandler)
	router.Run(":8080")
}

type patientResponse struct {
	Id    string `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}

func routerHandler(c *gin.Context) {
	patient := patientResponse{
		Id:    "2018-0001",
		Fname: "panumars",
		Lname: "seanto",
		Age:   23,
	}
	c.JSON(201, patient)
}
