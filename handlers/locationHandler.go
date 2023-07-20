package handlers

import (
	"fmt"
	services "groupie-tracker/Services"
	"html/template"
	"net/http"
	"strings"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ids := r.FormValue("id")

		locationsData := services.LocationData(ids)
		location := locationsData.Locations
		fmt.Println(location)
		if locationsData.ID == 0 {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "")
			return
		}

		parse, err := template.ParseFiles("templates/search.html", "templates/locations.html", "templates/base.html")
		if err != nil {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "")
			return
		}

		var locations []struct {
			Town    string
			Country string
		}
		for _, loc := range location {
			parts := strings.Split(loc, "-")
			locations = append(locations, struct {
				Town    string
				Country string
			}{
				Town:    parts[0],
				Country: parts[1],
			})
		}
		data := struct {
			Title        string
			TemplateName string
			Locations    []struct {
				Town    string
				Country string
			}
		}{
			Title:        "Les Artistes du 22eme siecle",
			TemplateName: "location.html",
			Locations:    locations,
		}

		err = parse.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "")
			return
		}
	} else {
		ErrorHandlerHelp(w, r, "not Allowed error: 405", http.StatusMethodNotAllowed, "")
		return
	}
}
