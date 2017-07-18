package main

import (
	"encoding/json"
	"net/http"
)

// VersionData struct for version
// swagger:response VersionData
type VersionData struct {
	Version   int     `required json:"version" description:"api player version"`
	GOVersion float32 `json:"go_version" description:"golang version"`
}

func (env *Env) VersionAPI(w http.ResponseWriter, r *http.Request) {

	// swagger:route GET /version tools getVersion
	//
	// get version and other.
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
	//       200: VersionData
	//       500: ServerError

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	version := VersionData{Version: 1}

	if err := json.NewEncoder(w).Encode(version); err != nil {
		panic(err)
	}
}
