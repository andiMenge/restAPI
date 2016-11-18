package main

import (
	"net/http"
)

const port = ":1337"

func main() {

	router := NewRouter()
	InfoLog.Printf("starting webserver on port %s", port)
	ErrLog.Fatal(http.ListenAndServe(port, router))
}
