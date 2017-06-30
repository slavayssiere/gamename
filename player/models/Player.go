package models

import "log"
import "gopkg.in/mgo.v2"

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
