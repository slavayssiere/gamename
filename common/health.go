package common

import (
	"encoding/json"
	"net/http"
)

type StHealthCheck struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	record := StHealthCheck{Status: "OK", Code: 200}
	if err := json.NewEncoder(w).Encode(record); err != nil {
		panic(err)
	}
}
