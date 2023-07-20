package services

import "fmt"

const ArtistAPIURL = "https://groupietrackers.herokuapp.com/api/artists"

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}


func ArtistData() []Artist {

	var artist []Artist
	err := GetData(ArtistAPIURL, &artist)
	if err != nil {
		fmt.Println("Erreur lors de la recuperation des donnees:", err)
		return nil
	}

	return artist
}

func SearchArtistData(name string) Artist {
	var artist []Artist
	var artistnil Artist

	err := GetData(ArtistAPIURL, &artist)
	if err != nil {
		fmt.Println("Erreur lors de la recuperation des donnees:", err)
		return artistnil
	}
	for i := 0; i < len(artist); i++ {
		if artist[i].Name == name {
			return artist[i]
		}
	}

	return artistnil
}

func IsStringInList(str string, list []string) bool {
	for _, item := range list {
		if str == item {
			return true
		}
	}
	return false
}
