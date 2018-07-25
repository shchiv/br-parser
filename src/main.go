package main

import (
	"github.com/br-parser/src/router"
	"github.com/oschwald/geoip2-golang"
	"github.com/pkg/errors"
	"log"
)

func main() {
	reader, err := createReader()
	if err != nil {
		//TODO handle an error in another way
		log.Fatalf("Can't load country db %s", err.Error())
	}
	defer reader.Close()

	server := router.NewServer(reader)
	log.Printf("Starting server ...")
	server.Start()
}

func createReader() (*geoip2.Reader, error) {
	if reader, err := geoip2.Open("GeoLite2-Country.mmdb"); err != nil {
		return nil, err
	} else if reader == nil {
		return nil, errors.New("GeoIP reader is nil")
	} else {
		return reader, nil
	}
}
