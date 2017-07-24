package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/slavayssiere/gamename/common"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Player is a player
// swagger:response Player
type Player struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
}

// PlayerController to manage mgo
type PlayerController struct {
	session *mgo.Session
}

// NewPlayerController test
func NewPlayerController(s *mgo.Session) *PlayerController {
	return &PlayerController{s}
}

func (pc PlayerController) playerCreate(w http.ResponseWriter, r *http.Request) {
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

	ds := pc.session.Copy()
	defer ds.Close()

	tmpuser := context.Get(r, common.AuthUser)
	user := tmpuser.(common.GoogleAuth)

	player := Player{FirstName: user.GivenName, LastName: user.FamillyName}
	player.createPlayer(ds)

	if err := json.NewEncoder(w).Encode(player); err != nil {
		panic(err)
	}
}

func (player *Player) createPlayer(session *mgo.Session) {
	c := session.DB("player").C("player")

	player.searchPlayer(session)

	if player.ID == "" {
		log.Println("create new player")
		player.ID = bson.NewObjectId()
		err := c.Insert(player)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (player *Player) searchPlayer(session *mgo.Session) {
	c := session.DB("player").C("player")

	if err := c.Find(bson.M{"firstname": player.FirstName}).One(player); err != nil {
		log.Println(err)
	}
}
