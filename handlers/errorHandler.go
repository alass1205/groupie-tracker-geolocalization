package handlers

import (
	"html/template"
	"net/http"
)

func ErrorHandlerHelp(w http.ResponseWriter, r *http.Request, erreur string, code int, img string) {

	parse, err := template.ParseFiles("templates/search.html", "templates/filter.html", "templates/geolocalisation.html", "templates/error.html", "templates/base.html")
	if err != nil {
		ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
		return
	}

	// Définition des données à utiliser dans le template
	data := struct {
		Title        string
		ErrorMessage string
		Img          string
		TemplateName string
	}{
		Title:        "Les Artistes du 20eme siecle",
		ErrorMessage: erreur,
		Img:          img,
		TemplateName: "error.html",
	}

	if code != 0 {
		// Vérifier si le code de statut a déjà été écrit
		if w.Header().Get("Content-Type") == "" {
			w.WriteHeader(code)
		}
	}

	err = parse.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, "Erreur lors du rendu du template", http.StatusInternalServerError)
		return
	}
}
