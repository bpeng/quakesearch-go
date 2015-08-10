# GeoNet Quake Search

Provide web service to search the GeoNet quake catalog

## Development
Application is developed in GO

### Database
This application searches data on the GeoNet wfs database (will be changed to GeoNet HAZ database)

## Web service api

Restful web service for search quakes in difference format( geojson, gml, kml, csv)
query parameters: date, location, depth/magnitude

## Interactive web interface
* use interactive map to define search area
* update coordinates by map extent when "Map Extent" selected as default
* allows building search query as well as showing search results on interactive map
* number of quakes to show on map limited to 2000.
* output format currently (query builder): geojson, gml, kml, csv.
* result as url(s) for intended data, also button to download data from browser.
* the maximum number of quakes for each request is limited to 20,000 (to prevent server crash), beyond that multiple requests are suggested.

## Test

Build/run docker image for test database

```
cd test_db
sudo docker build  -t 'quakesearch_db' .
sudo docker run --name quakesearch_db -p 5432:5432  -i -d -t quakesearch_db
```

Run test
```
godep go test
```

## Build / Deployment


### Configuration
config parameters can be specified in ```quakesearch.json```  which can be copied to ```/etc/sysconfig/ ```

config parameters can also be specified as environment variables which will overwrite the above config file:
```
export QUAKESEARCH_DATABASE_HOST=localhost
export QUAKESEARCH_DATABASE_NAME=hazard
export QUAKESEARCH_DATABASE_USER=hazard_r
export QUAKESEARCH_DATABASE_PASSWORD=####
export QUAKESEARCH_WEB_SERVER_PORT=8080
```

### GO
```
godep go build
./quakesearch
```

### Docker
```
sudo docker build  -t 'quakesearch_go' .
sudo docker run --name quakesearch_go -p 8080:8080 -e "QUAKESEARCH_WEB_SERVER_PORT=8080" -e "QUAKESEARCH_DATABASE_HOST=localhost"  -e "QUAKESEARCH_DATABASE_NAME=hazard" -e "QUAKESEARCH_DATABASE_USER=####" -e "QUAKESEARCH_DATABASE_PASSWORD=####" -i -d -t quakesearch_go
sudo docker stop/start quakesearch_go
```

