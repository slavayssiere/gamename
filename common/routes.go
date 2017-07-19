package common

import (
	"net/http"
)

// Route define a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slices of Route
type Routes []Route

var routes = Routes{
	Route{
		"Health",
		"GET",
		"/health",
		healthCheck,
	},
}
