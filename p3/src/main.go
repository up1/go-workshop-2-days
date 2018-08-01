package main

import (
	apiLibrary "api"
	"fmt"
	"net/http"
	"service"

	mgo "gopkg.in/mgo.v2"
)

const url = "mongodb://localhost:27017"

func main() {
	DBSession, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("Cannot connect database ", err.Error())
		return
	}
	patientService := service.PatientService{DBSession: DBSession}
	api := apiLibrary.Api{PatientService: patientService}
	http.HandleFunc("/v1/patients", api.CreatePatientHandler)

	http.ListenAndServe(":3000", nil)

}
