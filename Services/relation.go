package services

const RelationAPIURL = "https://groupietrackers.herokuapp.com/api/relation/"

type DatesLocations map[string][]string

type Relation struct {
	ID             int            `json:"id"`
	DatesLocations DatesLocations `json:"datesLocations"`
}

func RelationData(id string) Relation {
	var relation Relation

	err := GetData(RelationAPIURL+id, &relation)
	if err != nil {
		return relation
	}
	return relation
}
