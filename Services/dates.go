package services

import "fmt"

const DatesAPIURL = "https://groupietrackers.herokuapp.com/api/dates/"

type Dates struct {
	ID   int      `json:"id"`
	Date []string `json:"dates"`
}

func DateData(id string) Dates {
	var date Dates

	err := GetData(DatesAPIURL+id, &date)
	if err != nil {
		fmt.Println("Erreur lors de la recuperation des donnees:", err)

	}
	return date
}
