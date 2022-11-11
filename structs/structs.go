package structs

type University struct {
	Country      string   `json:"country"`
	Name         string   `json:"name"`
	AlphaTwoCode string   `json:"alpha_two_code"`
	WebPages     []string `json:"web_pages"`
}

type Country struct {
	Names     Name              `json:"name"`
	Languages map[string]string `json:"languages"`
	Maps      Map               `json:"maps"`
	Cca2      string            `json:"cca2"`
}

type Map struct {
	GoogleMaps     string `json:"googleMaps"`
	OpenStreetMaps string `json:"openStreetMaps"`
}

type Name struct {
	Common      string            `json:"common"`
	Official    string            `json:"official"`
	NativeNames map[string]string `json:"native_names"`
}

type Bordering struct {
	Borders []string `json:"borders"`
	Cca3    string   `json:"cca3"`
}

type CombinedUniversityAndCountryStruct struct {
	Name         string            `json:"name"`
	Country      string            `json:"country"`
	AlphaTwoCode string            `json:"isocode"`
	WebPages     []string          `json:"web_pages"`
	Languages    map[string]string `json:"languages"`
	Map          string            `json:"map"`
}

type Diagnostics struct {
	Universitiesapi int    `json:"universitiesapi"`
	Countriesapi    int    `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          string `json:"uptime"`
}
