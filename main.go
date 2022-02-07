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

type AllData struct {
	URLs    URLs
	Artists []Artist
}

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

	allData := AllData{
		URLs:    urls,
		Artists: artists,
	}

	page.Execute(w, allData)
}

func artist(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/artist.html"))
	id := r.URL.Query().Get("id") // Parse "?id=""
	source := strings.Replace("https://groupietrackers.herokuapp.com/api/artists/:id", ":id", id, 1)

	// GET artist/:id
	response, _ := http.Get(source)
	data, _ := ioutil.ReadAll(response.Body)
	var artist Artist
	json.Unmarshal(data, &artist)

	page.Execute(w, artist)
}

// Lancement du serveur web

func main() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/artist", artist)

	http.ListenAndServe(":8080", nil)
}
