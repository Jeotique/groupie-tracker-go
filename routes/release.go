package routes

import (
	"fmt"
	"log"
	"net/http"
	"project/functions"
	"project/variables"
)

func ReleaseHandler(w http.ResponseWriter, r *http.Request) {
	artist, _ := functions.GetReleased(variables.AccessToken)
	fmt.Println(artist)
	if err := variables.Tmpl.ExecuteTemplate(w, "release", artist); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
		return
	}
}
