package handlers

import (
	/*Le package html/template est une bibliothèque de Go
		qui permet de générer des pages
		HTML dynamiques en utilisant des templates.
		Le package html/template est conçu pour générer des pages HTML en évitant
		les problèmes de sécurité liés à l'injection de code malveillant dans les templates.
		Il s'assure que toutes les valeurs insérées dans les templates
		sont correctement échappées pour prévenir les attaques XSS.

		405 : http.StatusMethodNotAllowed
	    400 : http.StatusBadRequest
	    404 : http.StatusNotFound
	    500 : http.StatusInternalServerError
	*/

	"fmt"
	services "groupie-tracker/Services"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var cachedArtist *[]services.Artist

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Utilisation du moteur de rendu de templates pour generer la reponse HTML
		parse, err := template.ParseFiles("templates/search.html", "templates/index.html", "templates/base.html")
		if err != nil {

			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}

		// Vérifie si le chemin de l'URL est différent de "/"
		if r.URL.Path != "/" && r.URL.Path != "/search" {

			ErrorHandlerHelp(w, r, "Page not found :\nError 404", http.StatusNotFound, "error404.png")
			return
		}
		artists := services.ArtistData()
		cachedArtist = &artists
		locationAll := services.LocationDataAll()
		if artists == nil {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}
		data := struct {
			Title        string
			Artists      []services.Artist
			TemplateName string
			LocationAll  services.LocationAll
		}{
			Title:        "Les Artistes du 20eme siecle",
			Artists:      artists,
			TemplateName: "index.html",
			LocationAll:  locationAll,
		}

		err = parse.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			fmt.Println(err)
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}
	} else if r.Method == http.MethodPost {

		if r.URL.Path != "/search" {
			fmt.Println("al")
			ErrorHandlerHelp(w, r, "Page not found :\nError 404", http.StatusNotFound, "error404.png")
			return
		}
		// Utilisation du moteur de rendu de templates pour generer la reponse HTML
		parse, err := template.ParseFiles("templates/search.html", "templates/index.html", "templates/base.html")
		if err != nil {
			fmt.Println(err)
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}

		Artists := services.ArtistData()
		locationAll := services.LocationDataAll()
		if Artists == nil {
			fmt.Println("la")
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}

		value := r.FormValue("value")
		NewArtists := []services.Artist{}
		returne := false

		if !strings.Contains(value, "_") {
			for _, artist := range Artists {
				if strings.EqualFold(value, artist.Name) ||
					strings.Contains(strings.ToLower(artist.Name), strings.ToLower(value)) ||
					value == artist.FirstAlbum ||
					strings.Contains(strconv.Itoa(artist.CreationDate), value) {
					NewArtists = append(NewArtists, artist)
					returne = true
				} else if IsExist(value, artist.Members) {
					NewArtists = append(NewArtists, artist)
					returne = true
				}
			}
		}

		tabId := map[int]bool{}
		for i, loc := range locationAll.Index {
			for _, add := range loc.Locations {
				if strings.Contains(strings.ToLower(add), strings.ToLower(value)) {
					if !tabId[i] {
						tabId[i] = true
						NewArtists = append(NewArtists, Artists[i])
						returne = true
					}
				}
			}
		}

		if !returne {
			fmt.Println("ici")
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}
		data := struct {
			Title        string
			Artists      []services.Artist
			TemplateName string
			LocationAll  services.LocationAll
		}{
			Title:        "Les Artistes du 20eme siecle",
			Artists:      NewArtists,
			TemplateName: "index.html",
			LocationAll:  locationAll,
		}

		err = parse.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			fmt.Println(err)
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}
	} else {
		fmt.Println("laaaa")
		ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
		return
	}
}

func IsExist(value string, members []string) bool {
	for _, member := range members {
		if strings.EqualFold(member, value) || strings.Contains(strings.ToLower(member), strings.ToLower(value)) {
			return true
		}
	}
	return false
}
