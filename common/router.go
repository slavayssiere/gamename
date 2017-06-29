package common

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewRouter add all routes
func NewRouter(routesAsk Routes) *mux.Router {
	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "uri_duration_seconds",
		Help: "Time to respond",
	}, []string{"uri"})

	router := mux.NewRouter().StrictSlash(true)

	routes = append(routes, routesAsk...)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name, histogram)

		fmt.Println(route)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.Methods("GET").Path("/metrics").Name("Metrics").Handler(promhttp.Handler())

	prometheus.Register(histogram)

	return router
}
