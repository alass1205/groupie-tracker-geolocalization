package handlers

import (
	"fmt"
	services "groupie-tracker/Services"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		artist := cachedArtist
		if artist == nil {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}
		ids := r.FormValue("id")
		id, err := strconv.Atoi(ids)

		if err != nil || (id < 1 && id > len(*artist)) {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}
		relation := services.RelationData(ids)

		if relation.ID == 0 {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}

		funcs := template.FuncMap{
			"split": strings.Split,
		}
		parse, err := template.New("artist.html").Funcs(funcs).ParseFiles("templates/search.html", "templates/filter.html", "templates/artist.html", "templates/base.html")
		if err != nil {
			fmt.Println("a")
			fmt.Println(err)
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}

		id = id - 1
		data := struct {
			TemplateName string
			ID           string
			Title        string
			Name         string
			Members      []string
			Image        string
			FirstAlbum   string
			CreationDate string
			Relations    []struct {
				Location          string
				Dates             []string
				RandomCoordinates struct {
					Lat float64
					Lng float64
				}
			}
		}{
			TemplateName: "artist.html",
			ID:           strconv.Itoa((*artist)[id].ID),
			Title:        "Les Artistes du 22eme siecle",
			Name:         (*artist)[id].Name,
			Members:      (*artist)[id].Members,
			Image:        (*artist)[id].Image,
			FirstAlbum:   (*artist)[id].FirstAlbum,
			CreationDate: strconv.Itoa((*artist)[id].CreationDate),
			Relations: []struct {
				Location          string
				Dates             []string
				RandomCoordinates struct {
					Lat float64
					Lng float64
				}
			}{},
		}

		for loc, dates := range relation.DatesLocations {
			coordinates, err := services.GetCoordinates(loc)
			if err != nil {
				fmt.Println(err)
				ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
				return
			}

			rel := struct {
				Location          string
				Dates             []string
				RandomCoordinates struct {
					Lat float64
					Lng float64
				}
			}{
				Location: loc,
				Dates:    dates,
				RandomCoordinates: struct {
					Lat float64
					Lng float64
				}{
					Lat: coordinates.Results[0].Locations[0].LatLng.Lat,
					Lng: coordinates.Results[0].Locations[0].LatLng.Lng,
				},
			}
			data.Relations = append(data.Relations, rel)
		}

		err = parse.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			fmt.Println("b alou	")
			fmt.Println(err)
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}
	} else {
		ErrorHandlerHelp(w, r, "Not allowed: 405", http.StatusMethodNotAllowed, "error500.png")
		return
	}
}
