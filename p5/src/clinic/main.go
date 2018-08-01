package main

import (
	"clinic/handler"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api/v1/patients", handler.CreateNewPatient)
	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
