package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome!")
}

func Foo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "bar")
}

func WriteRedis(w http.ResponseWriter, r *http.Request) {
	//get url query params
	redisKey := r.URL.Query().Get("key")
	redisValue := r.URL.Query().Get("value")
	if redisValue == "" || redisKey == "" {
		http.Error(w, "key or value not set", http.StatusBadRequest)
		return
	}

	//write params to redis
	err := redisSet(redisKey, redisValue)
	if err != nil {
		ErrLog.Printf("error writing to redis: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "write to redis successful")
}

func ReadRedis(w http.ResponseWriter, r *http.Request) {
	redisKey := r.URL.Query().Get("key")
	if redisKey == "" {
		http.Error(w, "key not set", http.StatusBadRequest)
		return
	}
	value, err := redisGet(redisKey)
	if err != nil {
		ErrLog.Printf("error reading from redis: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "value=%s", value)
}
