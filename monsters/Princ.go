package main

import (
	"encoding/json"
	"net/http"
)

// HelloWorld struct test
// swagger:response HelloWorld
type HelloWorld struct {
	Hello string `json:"hello"`
}

// Index function for slash
func index(w http.ResponseWriter, r *http.Request) {

	// swagger:route GET / tools getSlash
	//
	// get / and other.
	//
	// blabla line 1
	// BLABLA line 2
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https, ws, wss
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: HelloWorld
	//       500: ServerError

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	record := HelloWorld{Hello: "world"}

	if err := json.NewEncoder(w).Encode(record); err != nil {
		panic(err)
	}
}
