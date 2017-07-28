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
	Protected   bool
}

// Routes slices of Route
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Health",
		Method:      "GET",
		Pattern:     "/health",
		HandlerFunc: healthCheck,
		Protected:   false,
	},
}
