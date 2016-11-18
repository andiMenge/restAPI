package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var (
	InfoLog *log.Logger
	ErrLog  *log.Logger
)

func init() {
	// create new logger for Info and Err msg
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	ErrLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		InfoLog.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
