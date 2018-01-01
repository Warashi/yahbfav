package main

import (
	"encoding/xml"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

//go:generate go-assets-builder -s="/data" -o bindata.go data

var tmpl = template.Must(parseFile("/template.html"))

func Favorite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	resp, err := http.Get("http://b.hatena.ne.jp/" + ps.ByName("name") + "/favorite.rss")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var rss RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	tmpl.Execute(w, rss.Items)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.GET("/:name", Favorite)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
