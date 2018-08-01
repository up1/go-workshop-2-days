package api

import (
	"encoding/json"
	"net/http"
	"patient/model"
	"patient/service"
)

type Api struct {
	PatientService service.PatientService
}

func NewApi(patientService service.PatientService) Api {
	return Api{
		PatientService: patientService,
	}
}

func (api Api) CreatePatientHandler(writer http.ResponseWriter, request *http.Request) {
	var patient model.Patient
	err := json.NewDecoder(request.Body).Decode(&patient)
	if err != nil {
		http.Error(writer, "Not Found", http.StatusNotFound)
		return
	}
	newPatient, err := api.PatientService.InsertPatient(patient)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(writer).Encode(newPatient)
}
