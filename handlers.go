package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}
