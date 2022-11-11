package json_coder

/*
* 	This class is purely made for refactoring and making it so that each .go file has a
* 	specific purpose. This file ensures that arrays of University and Country gets decoded the same way.
 */

import (
	"assignment-1/structs"
	"encoding/json"
	"log"
	"net/http"
)

// DecodeUniversityInfo is a method that takes a http request and decodes the json body
func DecodeUniversityInfo(httpResponse *http.Response) []structs.University {
	decoder := json.NewDecoder(httpResponse.Body)
	var universities []structs.University

	if err := decoder.Decode(&universities); err != nil {
		log.Print(err, http.StatusNoContent)
	}

	return universities
}

// DecodeCountryInfo is a method that takes a http request and decodes the json body
func DecodeCountryInfo(httpResponse *http.Response) []structs.Country {
	decoder := json.NewDecoder(httpResponse.Body)
	var countries []structs.Country

	if err := decoder.Decode(&countries); err != nil {
		log.Print(err, http.StatusNoContent)
	}

	return countries
}

// DecodeCountryInfo is a method that takes a http request and decodes the json body
func DecodeBorderCountryInfo(httpResponse *http.Response) []structs.Bordering {
	decoder := json.NewDecoder(httpResponse.Body)
	var countries []structs.Bordering

	if err := decoder.Decode(&countries); err != nil {
		log.Print(err, http.StatusNoContent)
	}

	return countries
}
