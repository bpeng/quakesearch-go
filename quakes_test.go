package main

import (
	"encoding/json"
	"github.com/GeoNet/app/web"
	"github.com/GeoNet/app/web/webtest"
	"log"
	"net/http"
	"testing"
)

func TestQuakesCount(t *testing.T) {
	setup()
	defer teardown()
	//1. get all quakes
	c := webtest.Content{
		Accept: web.V1JSON,
		URI:    "/count?bbox=163.60840,-49.18170,182.98828,-32.28713",
	}
	b, err := c.Get(ts)
	if err != nil {
		t.Fatal(err)
	}
	var qc QuakesCount
	err = json.Unmarshal(b, &qc)
	if err != nil {
		log.Fatal(err)
	}
	if qc.Count != 3 {
		t.Errorf("Found wrong number of quakes: %d", qc.Count)
	}

	//2. get only one quake
	c = webtest.Content{
		Accept: web.V1JSON,
		URI:    "/count?bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2010-1-1T00:00:00&enddate=2015-1-1T00:00:00",
	}
	b, err = c.Get(ts)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(b, &qc)
	if err != nil {
		log.Fatal(err)
	}
	if qc.Count != 1 {
		t.Errorf("Found wrong number of quakes: %d", qc.Count)
	}

	//3. get 2 quakes
	c = webtest.Content{
		Accept: web.V1JSON,
		URI:    "/count?bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=5",
	}
	b, err = c.Get(ts)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(b, &qc)
	if err != nil {
		log.Fatal(err)
	}
	if qc.Count != 2 {
		t.Errorf("Found wrong number of quakes: %d", qc.Count)
	}
}

func TestQuakesGeoJson(t *testing.T) {
	setup()
	defer teardown()
	//1. get all quakes
	c := webtest.Content{
		Accept: web.V1GeoJSON,
		URI:    "/geojson?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713",
	}
	b, err := c.Get(ts)
	if err != nil {
		t.Fatal(err)
	}
	var f GeoJsonFeatureCollection
	err = json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal(err)
	}
	if len(f.Features) != 3 {
		t.Errorf("Found wrong number of features: %d", len(f.Features))
	}

	//2. get only one quake
	c = webtest.Content{
		Accept: web.V1GeoJSON,
		URI:    "/geojson?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2010-1-1T00:00:00&enddate=2015-1-1T00:00:00",
	}
	b, err = c.Get(ts)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal(err)
	}
	if len(f.Features) != 1 {
		t.Errorf("Found wrong number of features: %d", len(f.Features))
	}
	if f.Features[0].Properties.Publicid != "3366146" {
		t.Errorf("Found wrong publicid: %d", f.Features[0].Properties.Publicid)
	}

}

func TestRoutes(t *testing.T) {
	setup()
	defer teardown()

	//1 GeoJSON routes
	r := webtest.Route{
		Accept:     web.V1GeoJSON,
		Content:    web.V1GeoJSON,
		Cache:      web.MaxAge300,
		Surrogate:  web.MaxAge300,
		Response:   http.StatusOK,
		Vary:       "Accept",
		TestAccept: false,
	}
	r.Add("/geojson?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713")
	r.Add("/geojson?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2010-1-1T00:00:00&enddate=2015-1-1T00:00:00")
	r.Add("/geojson?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=3&maxmag=10")
	r.Add("/geojson?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&mindepth=10&maxdepth=200")
	r.Add("/geojson?limit=100&region=canterbury&minmag=3&maxmag=7&mindepth=1&maxdepth=200")
	r.Test(ts, t)

	//2. Count
	r = webtest.Route{
		Accept:     web.V1JSON,
		Content:    web.V1JSON,
		Cache:      web.MaxAge300,
		Surrogate:  web.MaxAge300,
		Response:   http.StatusOK,
		Vary:       "Accept",
		TestAccept: false,
	}
	r.Add("/count?bbox=163.60840,-49.18170,182.98828,-32.28713")
	r.Add("/count?bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2010-1-1T00:00:00&enddate=2015-1-1T00:00:00")
	r.Add("/count?bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=3&maxmag=10")
	r.Add("/count?bbox=163.60840,-49.18170,182.98828,-32.28713&mindepth=10&maxdepth=200")
	r.Add("/count?region=canterbury&minmag=3&maxmag=7&mindepth=1&maxdepth=200")
	r.Test(ts, t)

	//3 CSV routes
	r = webtest.Route{
		Accept:     web.V1CSV,
		Content:    web.V1CSV,
		Cache:      web.MaxAge300,
		Surrogate:  web.MaxAge300,
		Response:   http.StatusOK,
		Vary:       "Accept",
		TestAccept: false,
	}
	r.Add("/csv?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713")
	r.Add("/csv?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2010-1-1T00:00:00&enddate=2015-1-1T00:00:00")
	r.Add("/csv?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=3&maxmag=10")
	r.Add("/csv?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&mindepth=10&maxdepth=200")
	r.Add("/csv?limit=100&region=canterbury&minmag=3&maxmag=7&mindepth=1&maxdepth=200")
	r.Test(ts, t)

	//4 GML routes
	r = webtest.Route{
		Accept:     CONTENT_TYPE_XML,
		Content:    CONTENT_TYPE_XML,
		Cache:      web.MaxAge300,
		Surrogate:  web.MaxAge300,
		Response:   http.StatusOK,
		Vary:       "Accept",
		TestAccept: false,
	}
	r.Add("/gml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713")
	r.Add("/gml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2010-1-1T00:00:00&enddate=2015-1-1T00:00:00")
	r.Add("/gml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=3&maxmag=10")
	r.Add("/gml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&mindepth=10&maxdepth=200")
	r.Add("/gml?limit=100&region=canterbury&minmag=3&maxmag=7&mindepth=1&maxdepth=200")
	r.Test(ts, t)

	//5 KML routes
	r = webtest.Route{
		Accept:     CONTENT_TYPE_KML,
		Content:    CONTENT_TYPE_KML,
		Cache:      web.MaxAge300,
		Surrogate:  web.MaxAge300,
		Response:   http.StatusOK,
		Vary:       "Accept",
		TestAccept: false,
	}
	r.Add("/kml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713")
	r.Add("/kml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2010-1-1T00:00:00&enddate=2015-1-1T00:00:00")
	r.Add("/kml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=3&maxmag=10")
	r.Add("/kml?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&mindepth=10&maxdepth=200")
	r.Add("/kml?limit=100&region=canterbury&minmag=3&maxmag=7&mindepth=1&maxdepth=200")
	r.Test(ts, t)
}

func TestGeoJSON(t *testing.T) {
	setup()
	defer teardown()

	// GeoJSON routes
	r := webtest.Route{
		Accept:     web.V1GeoJSON,
		Content:    web.V1GeoJSON,
		Cache:      web.MaxAge300,
		Surrogate:  web.MaxAge300,
		Response:   http.StatusOK,
		Vary:       "Accept",
		TestAccept: false,
	}
	r.Add("/geojson?limit=100&bbox=163.60840,-49.18170,182.98828,-32.28713&startdate=2000-1-1T00:00:00&enddate=2015-1-1T00:00:00")

	r.GeoJSON(ts, t)
}

type QuakesCount struct {
	Count int      `json:"count"`
	Dates []string `json:"dates"`
}
