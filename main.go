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

func main() {
	http.HandleFunc("/", homepage)

	http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/homepage.html"))

	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")
	data, _ := ioutil.ReadAll(response.Body)
	var object IndexData
	json.Unmarshal(data, &object)

	page.Execute(w, object)
}
