package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/context"
	"github.com/prometheus/client_golang/prometheus"
)

type key int

// AuthUser context key for user finding
const AuthUser key = 0

func LoggerMiddleware(inner http.Handler, name string, histogram *prometheus.HistogramVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		time := time.Since(start)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time,
		)

		histogram.WithLabelValues(r.RequestURI).Observe(time.Seconds())
	})
}

// GAuthMiddleware middleware for Google Auth
func GAuthMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var record GoogleAuth

		authorization := r.Header.Get("authorization")
		token := url.QueryEscape(authorization)

		url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", token)

		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
			log.Println(err)
		}

		if record.Email == "" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "{\"reason\":\"not authorized\"}")
		} else {
			context.Set(r, AuthUser, record)
			inner.ServeHTTP(w, r)
		}

	})
}
