package routes

import (
	"net/http"
	handle "scoringservice/handlers"
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
}
