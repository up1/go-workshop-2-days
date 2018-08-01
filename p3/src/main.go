package main

import (
	"net/http"
	"router"
)

const url = "mongodb://localhost:27017"

func main() {
	// server := router.SetupRouteEcho()
	// server.Start(":3000")

	// server := router.SetupRouteGin()
	// server.Run(":3000")
	mux := router.SetupRouteStandardLibrary()
	http.ListenAndServe(":3000", mux)
}
