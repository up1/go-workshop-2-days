package main

import (
	"handler"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/patient")
	{
		v1.GET("/", handler.CreatePatient)
	}
	router.Run()

}
