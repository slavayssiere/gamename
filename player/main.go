package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/slavayssiere/gamename/common"
)

// @APIVersion 1.0.0
// @APITitle Player API
// @APIDescription Player API for gamename
// @Contact sebastien.lavayssiere@gmail.com
// @TermsOfServiceUrl https://www.perdu.com
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

func main() {
	router := common.NewRouter(routes)
	common.ConsulManagement("player")

	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

var routes = common.Routes{
	common.Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	common.Route{
		"Version",
		"GET",
		"/version",
		VersionAPI,
	},
	common.Route{
		"IPs",
		"GET",
		"/ip",
		getIP,
	},
}
