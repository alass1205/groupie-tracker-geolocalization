package services

const LocationAPIURL = "https://groupietrackers.herokuapp.com/api/locations/"
const LocationAPIURLALL = "https://groupietrackers.herokuapp.com/api/locations"

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type LocationAll struct {
	Index []Location `json:"index"`
}

func LocationData(id string) Location {
	var location Location

	err := GetData(LocationAPIURL+id, &location)
	if err != nil {
		return location

	}
	return location

}



func LocationDataAll() LocationAll {
	var locationAll LocationAll


	err := GetData(LocationAPIURLALL, &locationAll)
	if err != nil {
		return locationAll
	}

	return locationAll
}
