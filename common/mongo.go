package common

import (
	"log"

	"gopkg.in/mgo.v2"
)

// ConnectDatabase blabla
func ConnectDatabase(clientvault *VaultClient, connectionString string) (*mgo.Session, error) {
	connectauth := GetSecret(clientvault, "secret/playerdb")

	log.Println(connectauth)

	mlgn := (connectauth["login"]).(string)
	mpwd := (connectauth["password"]).(string)

	connectionString = "mongodb://" + mlgn + ":" + mpwd + "@" + connectionString
	log.Println(connectionString)
	session, err := mgo.Dial(connectionString)
	if err != nil {
		log.Printf("erreur in connexion: %s", err)
		return nil, err
	}

	log.Println("Connected !")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
