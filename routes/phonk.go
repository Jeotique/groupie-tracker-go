package routes

import (
	"fmt"
	"log"
	"net/http"
	"project/functions"
	"project/variables"
)

func PhonkHandler(w http.ResponseWriter, r *http.Request) {
	id_artist := r.FormValue("id")
	fmt.Println(id_artist)

	artist, err := functions.GetPhonkPlaylists("37i9dQZF1DWWY64wDtewQt", variables.AccessToken)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error getting artist data: %v", err)
		return
	}
	fmt.Println(artist) // Vérifiez les données récupérées de l'API Spotify

	if err := variables.Tmpl.ExecuteTemplate(w, "phonk", artist); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}
