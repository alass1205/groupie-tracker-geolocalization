package handlers

import (
	"fmt"
	services "groupie-tracker/Services"
	"html/template"
	"net/http"
	"strings"
)

/** Cela peut etre utile lorsque vous avez plusieurs fichiers de modèle et que vous souhaitez effectuer
des opérations spécifiques sur un modèle particulier, telles que l'ajout de fonctions personnalisées**/
/**
template.FuncMap est un type qui associe des noms de fonctions aux implémentations de ces fonctions dans le
langage de template. Dans ce cas,
 la fonction "split" est associée à la fonction strings.Split
**/

func LocationHandlerAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		locationAll := services.LocationDataAll()

		funcs := template.FuncMap{
			"split": strings.Split,
		}
		parse, err := template.New("locationsAll.html").Funcs(funcs).ParseFiles("templates/search.html", "templates/filter.html", "templates/geolocalisation.html", "templates/locationsAll.html", "templates/base.html")

		if err != nil {
			fmt.Println(err)
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "")
			return
		}

		data := struct {
			Title        string
			LocationsAll []services.Location
			TemplateName string
		}{
			Title:        "Les Artistes du 22eme siècle",
			LocationsAll: locationAll.Index,
			TemplateName: "locationAll.html",
		}

		err = parse.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			fmt.Println("b")
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "")

			return
		}
	} else {
		ErrorHandlerHelp(w, r, "not Allowed error: 405", http.StatusMethodNotAllowed, "")
		return
	}
}
