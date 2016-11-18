package main

import (
	"log"
	"net/http"
)

const port = ":1337"

func main() {

	router := NewRouter()
	log.Printf("starting webserver on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
