package handler

import (
	"assignment-1/client"
	"assignment-1/constants"
	"assignment-1/json_coder"
	"assignment-1/structs"
	"fmt"
	"net/http"
	"path"
	_ "path"
	"strconv"
	"strings"
)

// NeighbouringUniversities is supposed to show you the neighbouring universities based on the one you search for
func NeighbouringUniversities(w http.ResponseWriter, r *http.Request) {
	// Using a switch for error handling.
	switch r.Method {
	case http.MethodGet:
		handleCountryGetRequest(w, r)
	default:
		http.Error(w, "Method not supported. Currently only GET are supported.", http.StatusNotImplemented)
		return
	}

}

// handleCountryGetRequest gets the bordering countries' information.
// Using the cca3 code and getting the country information about the bordering countries.
// At last it encodes the combined list to json.
func handleCountryGetRequest(w http.ResponseWriter, r *http.Request) {
	// Using path.Split() to split path immediately following the final slash, separating
	// it into a directory and file name component.
	countryPath, nameToSend := path.Split(r.URL.Path)
	nameToSend = strings.ReplaceAll(path.Base(r.URL.Path), " ", "%20")
	// countryPath is the country which is written in the URL
	countryPath = path.Base(countryPath)
	var limitAsInt int

	// fmt.Sprint() formats the link to be a single string.
	// Combining the country API and the country you want the information about.
	link := fmt.Sprint(constants.COUNTRY_INFO + countryPath)

	// Gets response from the country API and decodes this to find the bordering countries.
	// So, get a country' information using the country API -> Decode the borders of the country we
	// get the information about.
	responseFromCountryAPI, err := client.GetResponseFromURL(link)
	if err != nil {
		http.Error(w, "No content was found using the url.", http.StatusNoContent)
	}
	borders := json_coder.DecodeBorderCountryInfo(responseFromCountryAPI)

	// makes it so that we also get universities from the country we searched for
	// uses a for-loop to get each of the bordering codes in countries like China which
	// is linked to multiple countries.
	var borderArr []string
	for i := range borders {
		if borders != nil {
			borderArr = borders[i].Borders
			borderArr = append(borderArr, borders[i].Cca3)
		} else {
			http.Error(w, "No bordering countries found, only showing results from country inputted", http.StatusNoContent)
		}
	}

	// Uses the fmt.Spring() to format and strings.Join() to use the cca3 codes gotten from the borders,
	// and create a string separating them using a comma.
	// Reason being that in the restcountries API it separates borders using a comma.
	link = fmt.Sprint(constants.COUNTRY_INFO_USING_CCA2 + strings.Join(borderArr[:], ","))

	// Getting a response from this new link and returns a json response.
	// Decoding this countryinfo, getting all information the countries which was sent in.
	responseFromCountryAPI, err = client.GetResponseFromURL(link)
	if err != nil {
		http.Error(w, "No content was found using the url.", http.StatusNoContent)
	}
	countries := json_coder.DecodeCountryInfo(responseFromCountryAPI)

	// Getting the university information based on the name of the university.
	// These universities are decoded in the method UniversityInformationBasedOnName()
	var unis = UniversityInformationBasedOnName(nameToSend)
	// checking to see if the length of unis is 0 or nil
	client.NothingMatches(w, unis)

	// This encodes the combined data, appending the neighbouring countries and unis using
	// GetCompletedList()
	completedList := GetCompletedList(countries, unis)

	// Limit parameter:
	// Separating the url with the = sign to say that I want a number put in after the separator.
	// If no number is inputted, return the entire length of the list as limit.
	if len(r.URL.RawQuery) != 0 {
		parseRawQuery := strings.Split(r.URL.RawQuery, "=")
		limitAsString := parseRawQuery[1]
		limitAsInt, _ = strconv.Atoi(limitAsString)
	} else {
		limitAsInt = len(completedList)
	}
	completedList = limiter(completedList, limitAsInt)

	json_coder.PrettyPrint(w, completedList)
}

// GetCompletedList retrieves the universities from the bordering countries
// using two for-each loops, looping through the universities and countries and if the
// cca2 code is the same it will append() the new struct variable and add the info
// from the universities and countries and returns this appended slice.
func GetCompletedList(countries []structs.Country, university []structs.University) []structs.CombinedUniversityAndCountryStruct {
	combinedUniversityAndCountryStructs := make([]structs.CombinedUniversityAndCountryStruct, 0)
	for _, uni := range university {
		for _, country := range countries {
			if country.Cca2 == uni.AlphaTwoCode {
				combinedUniversityAndCountryStructs = append(combinedUniversityAndCountryStructs, CombineStructs(uni, country))
			}
		}
	}
	return combinedUniversityAndCountryStructs
}

func limiter(combinedStruct []structs.CombinedUniversityAndCountryStruct, limit int) []structs.CombinedUniversityAndCountryStruct {
	if limit >= len(combinedStruct) {
		return combinedStruct
	}
	if limit < len(combinedStruct) {
		// only return limit elements
		combinedStruct = append(combinedStruct[:limit], combinedStruct[len(combinedStruct):]...)
	}
	return combinedStruct
}
