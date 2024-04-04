package routes

import (
	"fmt"
	"log"
	"net/http"
	"project/functions"
	"project/variables"
)

func RapfrHandler(w http.ResponseWriter, r *http.Request) {
	id_artist := r.FormValue("id")
	fmt.Println(id_artist)

	artist, err := functions.GetRapPlaylists("37i9dQZF1DX1X23oiQRTB5", variables.AccessToken)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting artist data: %v", err)
		return
	}
	fmt.Println(artist) // Vérifiez les données récupérées de l'API Spotify

	if err := variables.Tmpl.ExecuteTemplate(w, "rapfr", artist); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}
