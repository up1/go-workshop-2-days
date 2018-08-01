package main

import (
	"handler"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/api/v1/patients", handler.CreateNewPatient)

	router.Run(":8080")

}
