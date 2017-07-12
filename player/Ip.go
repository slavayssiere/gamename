package main

// @SubApi version [/ip]
// @SubApi allow you to get ip of container [/ip]

import (
	"encoding/json"
	"net/http"

	"github.com/slavayssiere/gamename/common"
)

type IPData struct {
	IP            string `json:"ip"`
	ConnectPlayer string `json:"ip_player"`
}

// GetIP function to display IP
// @Title Get IP Information
// @Description Get IP Information
// @Accept json
// @Success 200 {object} string &quot;Success&quot;
// @Failure 401 {object} string &quot;Access denied&quot;
// @Failure 404 {object} string &quot;Not Found&quot;
// @Resource /ip
func (env *Env) GetIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	ipdata := IPData{
		IP:            common.GetOutboundIP(),
		ConnectPlayer: common.GetIpForService("player"),
	}

	if err := json.NewEncoder(w).Encode(ipdata); err != nil {
		panic(err)
	}
}
