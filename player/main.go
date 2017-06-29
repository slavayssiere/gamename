package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/slavayssiere/gamename/common"
)

func main() {
	router := common.NewRouter(routes)
	common.ConsulManagement("player")

	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
