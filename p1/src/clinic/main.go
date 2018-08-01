package main

import (
	"clinic/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("api/v1/patient", router.RouterCreate)
	r.Run(":8080")
}
