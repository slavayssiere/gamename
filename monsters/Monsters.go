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

// Monster is a player
// swagger:response Monster
type Monster struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
}

// MonsterController to manage mgo
type MonsterController struct {
	session *mgo.Session
}

// NewMonstersController test
func NewMonstersController(s *mgo.Session) *MonsterController {
	return &MonsterController{s}
}

func (pc MonsterController) monsterCreate(w http.ResponseWriter, r *http.Request) {
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

	player := Monster{FirstName: user.GivenName, LastName: user.FamillyName}
	player.createMonster(ds)

	if err := json.NewEncoder(w).Encode(player); err != nil {
		panic(err)
	}
}

func (monster *Monster) createMonster(session *mgo.Session) {
	c := session.DB("monster").C("monster")

	monster.searchMonster(session)

	if monster.ID == "" {
		log.Println("create new monster")
		monster.ID = bson.NewObjectId()
		err := c.Insert(monster)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (monster *Monster) searchMonster(session *mgo.Session) {
	c := session.DB("player").C("player")

	if err := c.Find(bson.M{"firstname": monster.FirstName}).One(monster); err != nil {
		log.Println(err)
	}
}
