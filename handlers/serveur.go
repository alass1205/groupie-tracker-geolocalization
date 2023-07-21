package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	colorGreen := "\033[32m" // Mise en place de couleur pour la lisibilité dans le terminal
	colorBlue := "\033[34m"
	colorYellow := "\033[33m"

	fmt.Println(string(colorBlue), "[SERVER_INFO] : Starting local Server...")
	css := http.FileServer(http.Dir("./static"))
	http.Handle("/styles/", http.StripPrefix("/styles/", css))

	// Ajouter la route pour le gestionnaire HomeHandler
	http.Handle("/", http.HandlerFunc(HomeHandler))
	http.Handle("/artist", http.HandlerFunc(ArtistHandler))
	http.Handle("/artistlocation", http.HandlerFunc(LocationHandler))
	http.Handle("/locations", http.HandlerFunc(LocationHandlerAll))
	http.Handle("/artistdate", http.HandlerFunc(DateHandler))
	http.Handle("/search", http.HandlerFunc(HomeHandler))

	// Démarrage du serveur HTTP
	fmt.Println(string(colorGreen), "[SERVER_READY] : on http://localhost:8080 ✅ ") // Mise en place de l'URL pour l'utilisateur
	fmt.Println(string(colorYellow), "[SERVER_INFO] : To stop the program : Ctrl + c \033[00m")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
