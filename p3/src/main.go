package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/v1/patients", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
}
