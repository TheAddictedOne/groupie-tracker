package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

// Définitions des structures

type URLs struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Router pour intercepter toutes les requêtes HTTP

func router(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	switch parts[1] {
	// /artist/:id
	// parts[0] => ""
	// parts[1] => artist
	// parts[2] => :id
	case "artist":
		artist(w, parts[2])

	default:
		homepage(w, r)
	}
}

// Route par défault

func homepage(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/homepage.html"))

	// GET index
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")
	data, _ := ioutil.ReadAll(response.Body)
	var urls URLs
	json.Unmarshal(data, &urls)

	// GET artists
	response, _ = http.Get(urls.Artists)
	data, _ = ioutil.ReadAll(response.Body)
	var artists []Artist
	json.Unmarshal(data, &artists)

	// Regroupement des données dans un map
	m := map[string]interface{}{
		"URLs":    urls,
		"Artists": artists,
	}

	page.Execute(w, m)
}

// Route /artist/:id

func artist(w http.ResponseWriter, id string) {
	page := template.Must(template.ParseFiles("views/artist.html"))
	source := strings.Replace("https://groupietrackers.herokuapp.com/api/artists/:id", ":id", id, 1)

	// GET artist/:id
	response, _ := http.Get(source)
	body, _ := ioutil.ReadAll(response.Body)
	var artist Artist
	json.Unmarshal(body, &artist)

	// GET relations de l'artist courant
	responseRelations, _ := http.Get(artist.Relations)
	bodyRelations, _ := ioutil.ReadAll(responseRelations.Body)
	var relation Relation
	json.Unmarshal(bodyRelations, &relation)

	// Regroupement des données dans un map
	m := map[string]interface{}{
		"Artist":   artist,
		"Relation": relation,
	}

	page.Execute(w, m)
}

// Lancement du serveur web

func main() {
	http.HandleFunc("/", router)

	http.ListenAndServe(":8080", nil)
}
