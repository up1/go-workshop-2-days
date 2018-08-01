package main

import (
	"api"
	"net/http"
)

func main() {

	http.HandleFunc("/createnewpatient", api.PostInfomationPatientID)
	http.ListenAndServe(":80", nil)

}
