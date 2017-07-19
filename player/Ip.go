package main

import (
	"encoding/json"
	"net/http"

	"github.com/slavayssiere/gamename/common"
)

// IPData struct for ip send
// swagger:response IPData
type IPData struct {

	// IP du container
	IP string `json:"ip"`
	// IP proposé par Consul
	ConnectPlayer string `json:"ip_player"`
}

// A GenericError is an error
// swagger:response genericError
type GenericError struct {
	// The error message
	// in: body
	Body struct {
		// The validation message
		//
		// Required: true
		Message string
		// An optional field name to which this validation applies
		FieldName string
	}
}

// A ServerError is an error huge
// swagger:response ServerError
type ServerError struct {
	// The error message
	// in: body
	Error struct {
		// The validation message
		//
		// Required: true
		Error error
	}
}

func getIP(w http.ResponseWriter, r *http.Request) {

	// swagger:route GET /ip tools getIp
	//
	// get ip and other.
	//
	// blabla line 1
	// BLABLA line 2
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Responses:
	//       default: genericError
	//       200: IPData
	//       500: ServerError

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	ipdata := IPData{
		IP:            common.GetOutboundIP(),
		ConnectPlayer: common.GetIPForService("player"),
	}

	if err := json.NewEncoder(w).Encode(ipdata); err != nil {
		panic(err)
	}
}
