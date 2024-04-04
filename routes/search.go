package routes

import (
	"fmt"
	"log"
	"net/http"
	"project/functions"
	"project/variables"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("search")
	fmt.Println(query)
	artist, _ := functions.SearchArtists(query, variables.AccessToken)
	fmt.Println(artist)
	if err := variables.Tmpl.ExecuteTemplate(w, "search", artist); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}
