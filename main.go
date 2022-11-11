package main

import (
	"assignment-1/constants"
	"assignment-1/handler"
	"log"
	"net/http"
	"os"
)

func handleRequests() {
	// Needed for Heroku on my computer
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Println("Port is not set, setting it to: " + "8080")
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {
	http.HandleFunc(constants.DEFAULT_PATH, handler.HomePage)
	http.HandleFunc(constants.INFO_PATH, handler.UniversityHandler)
	http.HandleFunc(constants.NEIGHBOUR_PATH, handler.NeighbouringUniversities)
	http.HandleFunc(constants.DIAG_PATH, handler.DiagnosticsHandler)
	handleRequests()
}
