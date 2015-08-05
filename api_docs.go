package main

import (
	"github.com/GeoNet/web"
	"github.com/GeoNet/web/api/apidoc"
)

const (
	WEB_KML = "application/vnd.google-earth.kml+xml"
	WEB_XML = "application/xml"
)

var docs = apidoc.Docs{
	Production:       config.WebServer.Production,
	APIHost:          config.WebServer.CNAME,
	Title:            `GeoNet QuakeSearch API`,
	Description:      `<p>The QuakeSearch API provides access the New Zealand earthquake catalogue, allows the user to search quakes using temporal, spatial, depth and magnitude constraints.</p>`,
	RepoURL:          `https://github.com/GeoNet/quakesearch-go`,
	StrictVersioning: false,
}

var exHost = "http://localhost:" + config.WebServer.Port

func init() {
	docs.AddEndpoint("geojson", &geojsonDoc)
	docs.AddEndpoint("gml", &gmlDoc)
	docs.AddEndpoint("kml", &kmlDoc)
	docs.AddEndpoint("csv", &csvDoc)
}

var geojsonDoc = apidoc.Endpoint{Title: "GeoJson",
	Description: `Get Quakes in GeoJson format.`,
	Queries: []*apidoc.Query{
		geojsonD,
	},
}

var geojsonD = &apidoc.Query{
	Accept:      web.V1GeoJSON,
	Title:       "GeoJson",
	Description: "Query quakes in GeoJson format for specified time, location, magnitude and depth.",
	Example:     "/geojson?bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=2&maxmag=7&mindepth=1&maxdepth=100&startdate=2015-7-5T22:00:00&enddate=2015-8-5T23:00:00",
	ExampleHost: exHost,
	URI:         "/geojson?bbox=(bbox)&minmag=(minmag)&maxmag=(maxmag)&mindepth=(mindepth)&maxdepth=(maxdepth)&startdate=(startdate)&enddate=(enddate)",
}

var gmlDoc = apidoc.Endpoint{Title: "GML",
	Description: `Get Quakes in GML format.`,
	Queries: []*apidoc.Query{
		gmlD,
	},
}

var gmlD = &apidoc.Query{
	Accept:      WEB_XML,
	Title:       "GML",
	Description: "Query quakes in GML format for specified time, location, magnitude and depth.",
	Example:     "/gml?bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=2&maxmag=7&mindepth=1&maxdepth=100&startdate=2015-7-5T22:00:00&enddate=2015-8-5T23:00:00",
	ExampleHost: exHost,
	URI:         "/gml?bbox=(bbox)&minmag=(minmag)&maxmag=(maxmag)&mindepth=(mindepth)&maxdepth=(maxdepth)&startdate=(startdate)&enddate=(enddate)",
}

var kmlDoc = apidoc.Endpoint{Title: "KML",
	Description: `Get Quakes in KML format.`,
	Queries: []*apidoc.Query{
		kmlD,
	},
}

var kmlD = &apidoc.Query{
	Accept:      WEB_KML,
	Title:       "KML",
	Description: "Query quakes in KML format for specified time, location, magnitude and depth.",
	Example:     "/kml?bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=2&maxmag=7&mindepth=1&maxdepth=100&startdate=2015-7-5T22:00:00&enddate=2015-8-5T23:00:00",
	ExampleHost: exHost,
	URI:         "/kml?bbox=(bbox)&minmag=(minmag)&maxmag=(maxmag)&mindepth=(mindepth)&maxdepth=(maxdepth)&startdate=(startdate)&enddate=(enddate)",
}

var csvDoc = apidoc.Endpoint{Title: "CSV",
	Description: `Get Quakes in CSV format.`,
	Queries: []*apidoc.Query{
		csvD,
	},
}

var csvD = &apidoc.Query{
	Accept:      web.V1CSV,
	Title:       "CSV",
	Description: "Query quakes in CSV format for specified time, location, magnitude and depth.",
	Example:     "/csv?bbox=163.60840,-49.18170,182.98828,-32.28713&minmag=2&maxmag=7&mindepth=1&maxdepth=100&startdate=2015-7-5T22:00:00&enddate=2015-8-5T23:00:00",
	ExampleHost: exHost,
	URI:         "/csv?bbox=(bbox)&minmag=(minmag)&maxmag=(maxmag)&mindepth=(mindepth)&maxdepth=(maxdepth)&startdate=(startdate)&enddate=(enddate)",
}
