package main

import (
	"database/sql"
	"github.com/GeoNet/cfg"
	"github.com/GeoNet/log/logentries"
	"github.com/GeoNet/web"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

//go:generate configer quakesearch.json
var (
	config = cfg.Load()
	db     *sql.DB
	header web.Header
)

func init() {
	logentries.Init(config.Logentries.Token)
	web.InitLibrato(config.Librato.User, config.Librato.Key, config.Librato.Source)
}

// main connects to the database, sets up request routing, and starts the http server.
func main() {
	var err error
	db, err = sql.Open("postgres", config.DataBase.Postgres())
	if err != nil {
		log.Fatalf("ERROR: problem with DB config: %s", err)
	}
	defer db.Close()

	db.SetMaxIdleConns(config.DataBase.MaxIdleConns)
	db.SetMaxOpenConns(config.DataBase.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		log.Println("Error: problem pinging DB - is it up and contactable?  500s will be served")
	}

	http.Handle("/", handler())
	log.Fatal(http.ListenAndServe(":"+config.WebServer.Port, nil))
}

// handler creates a mux and wraps it with default handlers.
func handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	return header.GetGzip(mux)
}
