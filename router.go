package main

import (
	"github.com/GeoNet/web/api/apidoc"
	"html/template"
	"net/http"
	"strings"
)

var indexTemp *template.Template

func init() {
	indexTemp = template.Must(template.ParseFiles("index.html"))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
}

func router(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.URL.Path == "/geojson":
		getQuakesGeoJson(w, r)
	case r.URL.Path == "/count":
		getQuakesCount(w, r)
	case r.URL.Path == "/csv":
		getQuakesCsv(w, r)
	case r.URL.Path == "/gml":
		getQuakesGml(w, r)
	case r.URL.Path == "/kml":
		getQuakesKml(w, r)

	case strings.HasPrefix(r.URL.Path, apidoc.Path):
		docs.Serve(w, r)
	default: //index page
		indexPage(w)
		//web.BadRequest(w, r, "Can't find a route for this request. Please refer to /api-docs")
	}
}

func indexPage(w http.ResponseWriter) {
	err := indexTemp.Execute(w, nil)
	if err != nil {
		http.Error(http.ResponseWriter(w), err.Error(), http.StatusInternalServerError)
	}
}
