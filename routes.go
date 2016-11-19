package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"foo",
		"GET",
		"/foo",
		Foo,
	},
	Route{
		"set",
		"GET",
		"/set",
		WriteRedis,
	},
	Route{
		"get",
		"GET",
		"/get",
		ReadRedis,
	},
}
