package main


// @SubApi Index [/]
// @SubApi allow you to test slash [/]

import (
	"encoding/json"
	"net/http"

	"github.com/slavayssiere/gamename/common"
)

// Index function for slash
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	record := "{'hello':'world'}"

	if err := json.NewEncoder(w).Encode(record); err != nil {
		panic(err)
	}
}

// func Connect(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	authorization := r.Header.Get("authorization")
// 	// QueryEscape escapes the phone string so
// 	// it can be safely placed inside a URL query
// 	safeAuth := url.QueryEscape(authorization)

// 	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%s", safeAuth)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer resp.Body.Close()

// 	var record GoogleAuth
// 	//var consultant Data.Consultant

// 	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
// 		log.Println(err)
// 	}
// 	// consultant.Email = record.Email
// 	// consultant.RepoFindConsultant()

// 	// if consultant.ID == 0 {
// 	// 	log.Println("First connexion of", record.Email)
// 	// 	temp := Data.Consultant{FirstName: record.GivenName, LastName: record.FamillyName, Email: record.Email}
// 	// 	temp.RepoCreateConsultant()
// 	// 	consultant = temp
// 	// 	log.Println("consultant: ", consultant)
// 	// }

// 	if err := json.NewEncoder(w).Encode(record); err != nil {
// 		panic(err)
// 	}
// }


// type GoogleAuth struct {
// 	Azp           string `json:"azp"`
// 	Aud           string `json:"aud"`
// 	Sub           string `json:"sub"`
// 	Hd            string `json:"hd"`
// 	Email         string `json:"email"`
// 	EmailVerified string `json:"email_verified"`
// 	Hash          string `json:"at_hash"`
// 	Iss           string `json:"iss"`
// 	Iat           string `json:"iat"`
// 	Exp           string `json:"exp"`
// 	Name          string `json:"name"`
// 	Picture       string `json:"picture"`
// 	GivenName     string `json:"given_name"`
// 	FamillyName   string `json:"family_name"`
// 	Locale        string `json:"locale"`
// 	Alg           string `json:"alg"`
// 	Kid           string `json:"kid"`
// 	Profil        string `json:"profil"`
}
