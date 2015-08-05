package main

import (
	"github.com/GeoNet/web"
	"github.com/GeoNet/web/api/apidoc"
	"net/http"
	"strings"
)

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
	default:
		web.BadRequest(w, r, "Can't find a route for this request. Please refer to /api-docs")
	}
}
