package main

import (
	"member"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/v1/patient", member.PatientHandle)
	router.Run(":3000")
}
