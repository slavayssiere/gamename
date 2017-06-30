package main

// @SubApi version [/version]
// @SubApi allow you to get version [/version]

import (
	"encoding/json"
	"net/http"
)

// VersionData struct for version
type VersionData struct {
	Version   int     `required json:"version" description:"api player version"`
	GOVersion float32 `json:"go_version" description:"golang version"`
}

// VersionAPI function to get version
// @Title Get Version Information
// @Description Get Version  Information
// @Accept json
// @Success 200 {object} string &quot;Success&quot;
// @Failure 401 {object} string &quot;Access denied&quot;
// @Failure 404 {object} string &quot;Not Found&quot;
// @Resource /version
func (env *Env) VersionAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	version := VersionData{Version: 1}

	if err := json.NewEncoder(w).Encode(version); err != nil {
		panic(err)
	}
}
