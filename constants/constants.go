package constants

const (
	// Paths for the urls
	DEFAULT_PATH   = "/"
	INFO_PATH      = "/unisearcher/v1/uniinfo/"
	NEIGHBOUR_PATH = "/unisearcher/v1/neighbourunis/"
	DIAG_PATH      = "/unisearcher/v1/diag/"

	// University url paths
	UNIVERSITY_SEARCH_INFO = "http://universities.hipolabs.com/search?name_contains="

	// Country url paths
	COUNTRY_INFO            = "https://restcountries.com/v3.1/name/"
	COUNTRY_INFO_USING_CCA2 = "https://restcountries.com/v3.1/alpha?codes="
	COUNTRY_FIELD_INFO      = "?fields=name,languages,maps"

	// Root url
	UNI_ROOT, COUNTRY_ROOT = "http://universities.hipolabs.com", "https://restcountries.com"
)
