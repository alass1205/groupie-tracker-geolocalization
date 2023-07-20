package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type coordinatesResponse struct {
	Results []struct {
		Locations []struct {
			LatLng struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"latLng"`
		} `json:"locations"`
	} `json:"results"`
}

//L'utilisation d'un pointeur est utile dans ce cas car la fonction GetCoordinates cree une nouvelle instance
//de la structure coordinatesResponse et la remplit avec les données de réponse de l'API.
// En retournant un pointeur vers cette instance, la fonction peut partager cette même instance de donnees
// avec la fonction appelante sans avoir à effectuer une copie supplementaire.

func GetCoordinates(location string) (*coordinatesResponse, error) {
	url := "https://www.mapquestapi.com/geocoding/v1/address?key=Pp1TLX6fSeXtgtVj9KS7uPVWhVx9bwum&location=" + location

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var coordinates coordinatesResponse
	err = json.Unmarshal(body, &coordinates)
	if err != nil {
		return nil, err
	}

	return &coordinates, nil
}
