package main

import (
	"clinic/router"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("api/v1/patient", router.RouterCreatePatient)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
