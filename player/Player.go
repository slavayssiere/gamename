package main

// @SubApi version [/ip]
// @SubApi allow you to get ip of container [/ip]

import (
	"encoding/json"
	"net/http"

	"github.com/slavayssiere/gamename/player/models"
)

// PlayerCreate function to create user
// @Title Create user
// @Description Get new user
// @Accept json
// @Success 204 {object} string &quot;Success&quot;
// @Failure 401 {object} string &quot;Access denied&quot;
// @Failure 404 {object} string &quot;Not Found&quot;
// @Resource /user
func (env *Env) PlayerCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	player1 := models.Player{FirstName: "sebastien", LastName: "lavayssiere"}
	player1.Create(env.session)

	if err := json.NewEncoder(w).Encode(player1); err != nil {
		panic(err)
	}
}
