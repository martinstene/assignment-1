package handler

import (
	"assignment-1/constants"
	"assignment-1/json_coder"
	"assignment-1/structs"
	"fmt"
	"log"
	"net/http"
	"time"
)

var startTime = time.Now()

// DiagnosticsHandler is used to showcase access to request content
func DiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	// Using a switch for error handling.
	switch r.Method {
	case http.MethodGet:
		json_coder.PrettyPrint(w, Combine())
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}

// Uses the time.Since() and checks how long it was since the API started
func uptime() time.Duration {
	return time.Since(startTime)
}

// Uses a GET call to the root link to the university API. Then return it's statuscode.
func getStatusUni() int {
	resp, err := http.Get(constants.UNI_ROOT)
	if err != nil {
		log.Print(err)
	}
	return resp.StatusCode
}

// Uses a GET call to the root link to the country API. Then return it's statuscode.
func getStatusCountry() int {
	resp, err := http.Get(constants.COUNTRY_ROOT)
	if err != nil {
		log.Print(err)
	}
	return resp.StatusCode
}

// Combine returns a struct of Diagnostics. Uses the methods created over and adds these.
func Combine() structs.Diagnostics {
	return structs.Diagnostics{
		Universitiesapi: getStatusUni(),
		Countriesapi:    getStatusCountry(),
		Version:         "v1",
		Uptime:          fmt.Sprint(uptime()),
	}
}
