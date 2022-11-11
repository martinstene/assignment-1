package handler

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/json_coder"
	"assignment-1/structs"
	"errors"
	"log"
	"net/http"
	"path"
	"strings"
)

// UniversityHandler acts as a switch method which only accepts GET requests.
func UniversityHandler(w http.ResponseWriter, r *http.Request) {
	// Using a switch for error handling.
	switch r.Method {
	case http.MethodGet:
		handleUniversityGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}
}

//handleUniversityGetRequest
func handleUniversityGetRequest(w http.ResponseWriter, r *http.Request) {
	//splits up the name, and gets the part of the path that contains the search, also replaces space with %20
	nameToSend := strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	//gets the information about the university with the searched name
	var uniReply = UniversityInformationBasedOnName(nameToSend)
	// checking to see if the length of uniReply is 0
	client.NothingMatches(w, uniReply)
	// Gets the uniReply and combines this with it's country counterpart from the
	// country API.
	completedList := CombineUniAndCountry(uniReply)
	json_coder.PrettyPrint(w, completedList)
}

// UniversityInformationBasedOnName uses a name of a university and returns an array
// of universities using a stringbuilder to create an url and decodes this response.
func UniversityInformationBasedOnName(nameOfUniversity string) []structs.University {
	// builds an url using strings.Builder
	originalURL := constants.UNIVERSITY_SEARCH_INFO
	var urlToSend strings.Builder
	// instead of + between the values, I used WriteToString
	urlToSend.WriteString(originalURL)
	urlToSend.WriteString(nameOfUniversity)

	response, err := http.Get(urlToSend.String())
	if err != nil {
		err.Error()
	}

	return json_coder.DecodeUniversityInfo(response)
}

func CombineUniAndCountry(universities []structs.University) []structs.CombinedUniversityAndCountryStruct {
	var combinedUniversityList []structs.CombinedUniversityAndCountryStruct
	// Uses map[string]interface{} to create a simple cache/value store that can be used locally
	// idea inspired from:
	//https://adityarama1210.medium.com/fast-golang-api-performance-with-in-memory-key-value-storing-cache-1b248c182bdb
	var countries = map[string]structs.Country{}
	// Uses a for-each loop to go through every university, using _ because I want
	// to loop through every university and have no need to seek out just 1 element.
	for _, universityList := range universities {
		nameOfCountry := universityList.Country
		country := structs.Country{}
		// creating cached values of countries so that it is added to the map
		// this shortens down the time by a lot, by reusing the same value from the map.
		if obtainedValue, mappedCountries := countries[nameOfCountry]; mappedCountries {
			country = obtainedValue
			// if the country isn't cached it will grab the new country and decode it using the common country name.
			// Then return an array of the combined struct
		} else if countryToBeReturned, err := DecodeCountry(nameOfCountry); err != nil {
			return []structs.CombinedUniversityAndCountryStruct{}
		} else {
			// sets the entry in the map to be the countryToBeReturned
			countries[nameOfCountry] = countryToBeReturned
		}
		country = countries[nameOfCountry]
		combinedUniversityList = append(combinedUniversityList, CombineStructs(universityList, country))
	}
	return combinedUniversityList
}

// CombineStructs Combines the university info with the country info and returns the combined struct.
func CombineStructs(uni structs.University, country structs.Country) structs.CombinedUniversityAndCountryStruct {
	return structs.CombinedUniversityAndCountryStruct{
		Name:         uni.Name,
		Country:      uni.Country,
		AlphaTwoCode: uni.AlphaTwoCode,
		WebPages:     uni.WebPages,
		Languages:    country.Languages,
		Map:          country.Maps.OpenStreetMaps,
	}
}

// DecodeCountry Gets the country based on the country name from the country api.
func DecodeCountry(nameOfCountry string) (structs.Country, error) {
	/* Using a strings.Builder to create an url that meets the country
	* APIs demands and then uses that to decode a country info
	* by getting back a http request from the country api.
	 */
	var urlToSend strings.Builder
	urlToSend.WriteString(constants.COUNTRY_INFO)
	urlToSend.WriteString(nameOfCountry)
	urlToSend.WriteString(constants.COUNTRY_FIELD_INFO)
	// how the url looks after: https://restcountries.com/v3.1/name/country?fields=name,languages,maps

	link, err := client.GetResponseFromURL(urlToSend.String())
	if err != nil {
		log.Print("No content found", http.StatusNoContent)
	}
	// uses the created url and gets a json response. this info is decoded after.
	country := json_coder.DecodeCountryInfo(link)

	// if no country is found it throws an error and return an empty struct
	if len(country) == 0 || country == nil {
		return structs.Country{}, errors.New("unable to retrieve the specified country")
	}

	// Returning the country, which is the first object grabbed, AKA the first search result.
	return country[0], nil
}
