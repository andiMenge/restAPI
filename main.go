package main

import (
	"net/http"
)

const port = ":1337"

func main() {
	err := initRedis()
	if err != nil {
		ErrLog.Fatalf("connection with redis failed: %s", err)
	}
	InfoLog.Printf("connection with redis successful")

	router := NewRouter()
	InfoLog.Printf("starting webserver on port %s", port)
	ErrLog.Fatal(http.ListenAndServe(port, router))
}
