package common

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type key int

const UserData key = 0

//Logger this is a logger
func Logger(inner http.Handler, name string, histogram *prometheus.HistogramVec) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// if strings.Contains(r.RequestURI, "/connect") {
		// 	inner.ServeHTTP(w, r)
		// } else {
		// 	authorization := r.Header.Get("authorization")
		// 	safeAuth := url.QueryEscape(authorization)

		// 	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", safeAuth)

		// 	resp, err := http.Get(url)
		// 	if err != nil {
		// 		log.Println(err)
		// 	}

		// 	defer resp.Body.Close()

		// 	var record GoogleAuth

		// 	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		// 		log.Println(err)
		// 	}

		// 	// if strings.EqualFold("wescale.fr", record.Hd) {
		// 	// 	var consultant Data.Consultant
		// 	// 	consultant.Email = record.Email
		// 	// 	consultant.RepoFindConsultant()
		// 	// 	context.Set(r, UserData, consultant)
		// 	// 	inner.ServeHTTP(w, r)
		// 	// } else {
		// 	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// 	// 	w.WriteHeader(http.StatusUnauthorized)
		// 	// 	fmt.Fprintf(w, "{\"reason\":\"not a wescaler\"}")
		// 	// }
		// }

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
