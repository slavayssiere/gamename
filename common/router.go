package common

import (
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewRouter add all routes
func NewRouter(routesAsk Routes, servicename string) *mux.Router {

	histogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: servicename + "_uri_duration_seconds",
		Help: "Time to respond",
	}, []string{servicename})

	// collector, err := zipkin.NewHTTPCollector(os.Getenv("JAEGER_HOST"))
	// if err != nil {
	// 	fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
	// 	os.Exit(-1)
	// }

	// hostPort := GetOutboundIP() + ":8080"
	// // create recorder.
	// recorder := zipkin.NewRecorder(collector, true, hostPort, servicename)

	// // create tracer.
	// tracer, err := zipkin.NewTracer(
	// 	recorder,
	// 	zipkin.ClientServerSameSpan(false), //same span can be set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	// 	zipkin.TraceID128Bit(true),         //make Tracer generate 128 bit traceID's for root spans.
	// )
	// if err != nil {
	// 	fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
	// 	os.Exit(-1)
	// }

	// // explicitly set our tracer to be the default tracer.
	// opentracing.InitGlobalTracer(tracer)

	router := mux.NewRouter().StrictSlash(true)
	routes = append(routes, routesAsk...)

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = LoggerMiddleware(handler, route.Name, histogram)
		//handler = middleware.FromHTTPRequest(tracer, route.Name)(handler)
		if route.Protected {
			handler = GAuthMiddleware(handler)
		}

		fmt.Println(route)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	fileHandler := http.StripPrefix("/swagger/", http.FileServer(http.Dir("./public/")))

	router.Methods("GET").Path("/metrics").Name("Metrics").Handler(promhttp.Handler())

	fileHandler = LoggerMiddleware(fileHandler, "Swagger", histogram)
	router.Methods("GET").PathPrefix("/").Name("Swagger").Handler(fileHandler)

	prometheus.Register(histogram)

	return router
}
