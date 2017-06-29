package main

// @SubApi version [/ip]
// @SubApi allow you to get ip of container [/ip]

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"
)

type IPData struct {
	IP string `json:"ip"`
}

// getIP function to display IP
// @Title Get IP Information
// @Description Get IP Information
// @Accept json
// @Success 200 {object} string &quot;Success&quot;
// @Failure 401 {object} string &quot;Access denied&quot;
// @Failure 404 {object} string &quot;Not Found&quot;
// @Resource /ip
func getIP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	ipdata := IPData{IP: GetOutboundIP()}

	if err := json.NewEncoder(w).Encode(ipdata); err != nil {
		panic(err)
	}
}

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")

	return localAddr[0:idx]
}
