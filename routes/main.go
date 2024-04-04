package routes

import (
	"fmt"
	"log"
	"net/http"
)

func InitRoutes() {

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/rapfr", RapfrHandler)
	http.HandleFunc("/phonk", PhonkHandler)
	http.HandleFunc("/pop", PopHandler)
	http.HandleFunc("/add-favorite", AddToFavorites)
	http.HandleFunc("/favorites", Favorite)
	http.HandleFunc("/release", ReleaseHandler)

	css := http.FileServer(http.Dir("./styles"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	// Démarrage du serveur HTTP
	fmt.Println("Serveur lancé sur http://localhost:8086")
	log.Fatal(http.ListenAndServe(":8086", nil))
}
