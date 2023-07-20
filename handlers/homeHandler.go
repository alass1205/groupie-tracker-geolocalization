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
	} else {
		ErrorHandlerHelp(w, r, "Page not found :\nError 404", http.StatusNotFound, "error404.png")
		return
	}
}
