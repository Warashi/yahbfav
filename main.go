package yahbfav

import (
	"encoding/xml"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

//go:generate go-assets-builder -p yahbfav -s "/data" -v Assets -o bindata.go data

var tmpl = template.Must(template.ParseFiles("data/template.html"))

func Favorite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := urlfetch.Client(appengine.NewContext(r))

	resp, err := c.Get("http://b.hatena.ne.jp/" + ps.ByName("name") + "/favorite.rss")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var rss RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, rss.Items)
}

func init() {
	router := httprouter.New()
	router.GET("/:name", Favorite)
	http.Handle("/", router)
}
