package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", homepage)

	http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/homepage.html"))
	page.Execute(w, nil)
	// lire un template avec des placeholders
	// Remplacer les placeholders avec les données récupérées de l'API
}
