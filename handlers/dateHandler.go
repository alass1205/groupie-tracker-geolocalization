package handlers

import (
	services "groupie-tracker/Services"
	"html/template"
	"net/http"
	"strings"
)

func DateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")

		dateData := services.DateData(id)
		date := dateData.Date
		if dateData.ID == 0 {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}

		parse, err := template.ParseFiles("templates/search.html", "templates/filter.html", "templates/dates.html", "templates/base.html")
		if err != nil {
			ErrorHandlerHelp(w, r, "internal error: 500", http.StatusInternalServerError, "error500.png")
			return
		}

		for i := 0; i < len(date); i++ {
			date[i] = strings.Replace(date[i], "*", "", 1)
		}
		data := struct {
			Title        string
			TemplateName string
			Date         []string
		}{
			Title:        "Les Artistes du 22eme siecle",
			TemplateName: "date.html",
			Date:         date,
		}

		err = parse.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			ErrorHandlerHelp(w, r, "not Allowed error: 405", http.StatusMethodNotAllowed, "")
			return
		}
	} else {
		ErrorHandlerHelp(w, r, "not Allowed error: 405", http.StatusMethodNotAllowed, "")
		return
	}
}
