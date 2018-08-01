package main

import (
	"log"
	"member"
	"net/http"
)

func main() {
	http.HandleFunc("/patient", member.GetPatient)
	http.HandleFunc("/patient/insert", member.Insert)
	log.Println("Server running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
