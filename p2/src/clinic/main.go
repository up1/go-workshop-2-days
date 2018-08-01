package main

import (
	"clinic/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/patient")
	{
		v1.POST("/", handler.CreatePatient)
	}
	router.Run()
}
