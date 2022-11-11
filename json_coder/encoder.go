package json_coder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
PrettyPrint gotten from 02-JSON-demo
Using an interface so that no extra method is needed.
Takes in a decoded json list and reformats it so that it looks
Pretty, hence the name. Uses a responseWriter to write it out on the API.
*/
func PrettyPrint(w http.ResponseWriter, completedList interface{}) {
	output, err := json.MarshalIndent(completedList, "", "  ")
	if err != nil {
		http.Error(w, "Error during pretty printing", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(output))
}
