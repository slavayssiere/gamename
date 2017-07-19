package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

// ConnectDatabase blabla
func ConnectDatabase(connectionString string) (*mgo.Session, error) {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Printf("erreur in connexion: %s", err)
		return nil, err
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
