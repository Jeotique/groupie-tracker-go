package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"project/variables"
)

func AddToFavorites(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id_artist := r.FormValue("id_artist")
	fmt.Println(id_artist)

	var favs variables.Favorites
	file, err := ioutil.ReadFile("favorites.json")
	if err != nil {
		// Gérer l'erreur ou créer le fichier s'il n'existe pas
	}
	json.Unmarshal(file, &favs)

	favs.IDs = append(favs.IDs, id_artist)

	// Écrire le fichier JSON mis à jour
	updatedFavs, err := json.Marshal(favs)
	if err != nil {
		// Gérer l'erreur
	}

	if err != nil && !os.IsNotExist(err) {
		http.Error(w, "Erreur lors de la lecture des favoris", http.StatusInternalServerError)
		return
	} else if os.IsNotExist(err) {
		favs = variables.Favorites{IDs: []string{}}
	}

	if err := ioutil.WriteFile("favorites.json", updatedFavs, 0644); err != nil {
		http.Error(w, "Erreur lors de la sauvegarde des favoris", http.StatusInternalServerError)
		return
	}

	// Redirection vers la page des favoris
	http.Redirect(w, r, "/favorites", http.StatusSeeOther)
}
