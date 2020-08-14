package routes

import (
	handle "log/handlers"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Welcome",
		"GET",
		"/",
		handle.Welcome,
	},
	Route{
		"Log",
		"POST",
		"/log",
		handle.Log,
	},
}
