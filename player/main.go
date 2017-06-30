package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/slavayssiere/gamename/common"
	"github.com/slavayssiere/gamename/player/models"
	mgo "gopkg.in/mgo.v2"
)

// @APIVersion 1.0.0
// @APITitle Player API
// @APIDescription Player API for gamename
// @Contact sebastien.lavayssiere@gmail.com
// @TermsOfServiceUrl https://www.perdu.com
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause

type Env struct {
	session *mgo.Session
}

func main() {
	db, err := models.ConnectDatabase("playerdb")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{session: db}

	var routes = common.Routes{
		common.Route{
			"Index",
			"GET",
			"/",
			env.Index,
		},
		common.Route{
			"Version",
			"GET",
			"/version",
			env.VersionAPI,
		},
		common.Route{
			"IPs",
			"GET",
			"/ip",
			env.GetIP,
		},
		common.Route{
			"PlayerCreate",
			"POST",
			"/player",
			env.PlayerCreate,
		},
	}

	router := common.NewRouter(routes)
	common.ConsulManagement("player")

	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
