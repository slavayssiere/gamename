// Package classification Monsters API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Sebastien Lavayssiere<sebastien.lavayssiere@wescale.fr> http://www.wescale.fr
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//        - email
//
//     SecurityDefinitions:
//     - api_key:
//          type: apiKey
//          name: authorization
//          in: header
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/slavayssiere/gamename/common"
	mgo "gopkg.in/mgo.v2"
)

type DataStore struct {
	session *mgo.Session
}

func main() {

	clientconsul := common.ConsulManagement("monsters")
	common.ListenService("player", clientconsul)

	appname := common.GetVariable("appconfig/appname")
	log.Printf("AppName: %s", appname)

	appversion := common.GetVariable("appconfig/appversion")
	log.Printf("AppVersion: %s", appversion)

	clientvault, err := common.VaultManagement()
	if err != nil {
		log.Fatalln(err)
	}

	addr := os.Getenv("MONGO_HOST")
	if len(addr) == 0 {
		addr = "127.0.0.1:27018"
	}
	db, err := ConnectDatabase(clientvault, addr)
	if err != nil {
		log.Panic(err)
	}
	mc := NewMonstersController(db)

	var routes = common.Routes{
		common.Route{
			Name:        "Index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: index,
			Protected:   false,
		},
		common.Route{
			Name:        "Version",
			Method:      "GET",
			Pattern:     "/version",
			HandlerFunc: versionAPI,
			Protected:   false,
		},
		common.Route{
			Name:        "ListMonsters",
			Method:      "GET",
			Pattern:     "/monsters",
			HandlerFunc: getMonsters,
			Protected:   false,
		},
		common.Route{
			Name:        "MonsterCreate",
			Method:      "POST",
			Pattern:     "/monster",
			HandlerFunc: mc.monsterCreate,
			Protected:   false,
		},
	}

	router := common.NewRouter(routes, "monsters")

	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
