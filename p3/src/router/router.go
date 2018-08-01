package router

import (
	"model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
)

func SetupRouteEcho() *echo.Echo {
	router := echo.New()
	router.POST("/api/v1/patients", func(context echo.Context) error {
		patientRequest := new(model.PatientRequest)
		if err := context.Bind(patientRequest); err != nil {
			return err
		}
		var patientResponse model.PatientResponse
		return context.JSON(http.StatusCreated, patientResponse)
	})
	return router
}

func SetupRouteGin() *gin.Engine {
	router := gin.Default()
	router.POST("/api/v1/patients", func(context *gin.Context) {
		patientRequest := new(model.PatientRequest)
		if err := context.Bind(patientRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		var patientResponse model.PatientResponse
		context.JSON(http.StatusCreated, patientResponse)
	})
	return router
}
