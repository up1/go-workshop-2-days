package main

import (
	"clinic/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	router.Run(":3000")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/patient")
	{
		v1.POST("/", handler.CreatePatient)
	}
	return router
}
