package main

// @SubApi version [/ip]
// @SubApi allow you to get ip of container [/ip]

import (
	"encoding/json"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
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

	player1 := Player{FirstName: "sebastien", LastName: "lavayssiere"}
	player1.Create(env.session)

	if err := json.NewEncoder(w).Encode(player1); err != nil {
		panic(err)
	}
}

type Player struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (player *Player) Create(session *mgo.Session) {
	c := session.DB("player").C("player")

	err := c.Insert(&player)

	if err != nil {
		log.Fatal(err)
	}
}
