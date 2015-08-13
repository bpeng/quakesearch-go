FROM quay.io/geonet/golang-godep:latest

COPY . /go/src/github.com/GeoNet/quakesearch

WORKDIR /go/src/github.com/GeoNet/quakesearch

RUN godep go install -a

EXPOSE 8080

CMD ["/go/bin/quakesearch"]
