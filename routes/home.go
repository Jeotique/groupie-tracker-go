package routes

import (
	"log"
	"net/http"
	"project/functions"
	"project/variables"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artist, _ := functions.GetAlbums("3IW7ScrzXmPvZhB27hmfgy", variables.AccessToken)
	if err := variables.Tmpl.ExecuteTemplate(w, "home", artist); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}
