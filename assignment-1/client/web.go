package client

import (
	"assignment-1/structs"
	"fmt"
	"net/http"
)

// NothingMatches method checks if the length of the array of universities is 0, if yes, throws an http error
func NothingMatches(w http.ResponseWriter, unis []structs.University) {
	if len(unis) == 0 || unis == nil {
		http.Error(w, "No University matching this search", http.StatusNoContent)
	}
	// else -> do nothing
}

// GetResponseFromURL method that takes an url and gets a json response from the webpage.
// Gotten from 05-REST-client in main.go
func GetResponseFromURL(url string) (*http.Response, error) {
	resp, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting the content type header
	resp.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}

	// Issue request
	res, err := client.Do(resp)
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	}

	return res, nil
}
