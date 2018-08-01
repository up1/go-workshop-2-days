package router

import (
	"model"
	"net/http"

	"github.com/labstack/echo"
)

func SetupRoute() {
	router := echo.New()
	router.POST("/api/v1/patients", func(context echo.Context) error {
		var patientRequest model.PatientRequest
		if err := context.Bind(patientRequest); err != nil {
			return err
		}
		var patientResponse model.PatientResponse
		return context.JSON(http.StatusCreated, patientResponse)
	})
}
