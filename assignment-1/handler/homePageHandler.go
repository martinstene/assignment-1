package handler

import (
	"assignment-1/constants"
	"fmt"
	"net/http"
)

const LINEBREAK = "\n"

// HomePage is a homepage just for the sake of having a homepage, showing the user commands on how to use the API
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "How to use the API: "+LINEBREAK)
	fmt.Fprintf(w, "Search for universities: "+constants.INFO_PATH+"{:partial_or_complete_university_name}/"+LINEBREAK)
	fmt.Fprintf(w, "Search for universities and their neighbouring universities sharing the same name: "+constants.NEIGHBOUR_PATH+"{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}"+LINEBREAK)
	fmt.Fprintf(w, "To show diagnostics use: "+constants.DIAG_PATH)
}
