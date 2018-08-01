package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/patient", hello)
	log.Println("Server running on port")
	http.ListenAndServe(":3000", nil)
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create new patient")
}
