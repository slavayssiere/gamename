package common

import (
	"encoding/json"
	"net/http"
)

// StatusHealthCheck data for health check
type StatusHealthCheck struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

// HealthCheck function to health check
// @Title Get health Information
// @Description Get health Information
// @Success 200 {object} string &quot;Success&quot;
// @Failure 500 {object} string &quot;Not Found&quot;
// @Resource /health
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	record := StatusHealthCheck{Status: "OK", Code: 200}
	if err := json.NewEncoder(w).Encode(record); err != nil {
		panic(err)
	}
}
