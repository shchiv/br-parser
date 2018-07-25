package main

import (
	"github.com/br-parser/src/router"
	"log"
)

func main() {
	reader, err := router.CreateReader()
	if err != nil {
		log.Fatalf("Can't load country db %s", err.Error())
	}
	defer reader.Close()

	server := router.NewServer(reader)
	log.Printf("Starting server ...")
	server.Start()
}
