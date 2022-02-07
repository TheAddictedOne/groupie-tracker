package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

type IndexData struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type ArtistsData []struct {
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
	IndexData   IndexData
	ArtistsData ArtistsData
}

func main() {
	http.HandleFunc("/", homepage)

	http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/homepage.html"))

	// GET index
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")
	data, _ := ioutil.ReadAll(response.Body)
	var indexData IndexData
	json.Unmarshal(data, &indexData)

	// GET artists
	response, _ = http.Get(indexData.Artists)
	data, _ = ioutil.ReadAll(response.Body)
	var artistsData ArtistsData
	json.Unmarshal(data, &artistsData)

	allData := AllData{
		IndexData:   indexData,
		ArtistsData: artistsData,
	}

	page.Execute(w, allData)
}
