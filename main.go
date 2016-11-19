package main

import (
	"net/http"
)

const port = ":1337"

func main() {
	err := initRedis()
	if err != nil {
		ErrLog.Fatalf("redis db not accessable: %s", err)
	}
	InfoLog.Printf("redis start successful")

	router := NewRouter()
	InfoLog.Printf("starting webserver on port %s", port)
	ErrLog.Fatal(http.ListenAndServe(port, router))
}
