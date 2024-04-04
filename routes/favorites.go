package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/functions"
	"project/variables"
)

func Favorite(w http.ResponseWriter, r *http.Request) {
	println("Favorites")
	var favs variables.Favorites
	file, err := ioutil.ReadFile("favorites.json")
	if err != nil {
		// Gérer l'erreur ou créer le fichier s'il n'existe pas
		favs = variables.Favorites{IDs: []string{}}
	}
	json.Unmarshal(file, &favs)

	var artists []variables.Album
	for _, id := range favs.IDs {
		println(id)
		artist, err := functions.GetAlbums(id, variables.AccessToken)
		if err != nil {
			// Gérer l'erreur
		}
		artists = append(artists, artist)
	}
	fmt.Println(artists)
	variables.Tmpl.ExecuteTemplate(w, "favorites", artists)
}
