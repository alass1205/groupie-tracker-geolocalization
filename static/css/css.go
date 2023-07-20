package css

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func CSS(w http.ResponseWriter, r *http.Request) {
	// Vérifie si l'URL correspond à un fichier CSS
	if strings.HasPrefix(r.URL.Path, "/static/") {
		// Récupère le chemin du fichier CSS
		cssPath := "/static/" + strings.TrimPrefix(r.URL.Path, "/static")
		// Lit le contenu du fichier CSS
		cssData, err := ioutil.ReadFile(cssPath)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture du fichier CSS", http.StatusInternalServerError)
			return
		}
		// Définit le type MIME pour les fichiers CSS
		w.Header().Set("Content-Type", "text/css")
		// Écrit le contenu du fichier CSS dans la réponse
		w.Write(cssData)
		return
	}
}
