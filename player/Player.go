package main

import (
	"encoding/json"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (env *Env) PlayerCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// swagger:route POST /player main createPlayer
	//
	// create player and tsss.
	//
	// blabla line 1
	// BLABLA line 2
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https, ws, wss
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: Player
	//       500: ServerError

	player1 := Player{FirstName: "aurelie", LastName: "samson"}
	player1.create(env.session)

	if err := json.NewEncoder(w).Encode(player1); err != nil {
		panic(err)
	}
}

// Player is a player
// swagger:response Player
type Player struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
}

func (player *Player) create(session *mgo.Session) {
	c := session.DB("player").C("player")

	err := c.Insert(&player)

	if err != nil {
		log.Fatal(err)
	}

	err = c.Find(bson.M{"firstname": "aurelie"}).One(&player)

	if err != nil {
		log.Fatal(err)
	}
}
