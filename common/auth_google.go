package common

// GoogleAuth struct for google authentication
type GoogleAuth struct {
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Hd            string `json:"hd"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Hash          string `json:"at_hash"`
	Iss           string `json:"iss"`
	Iat           string `json:"iat"`
	Exp           string `json:"exp"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamillyName   string `json:"family_name"`
	Locale        string `json:"locale"`
	Alg           string `json:"alg"`
	Kid           string `json:"kid"`
	Profil        string `json:"profil"`
}
