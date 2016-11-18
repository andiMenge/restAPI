package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}
