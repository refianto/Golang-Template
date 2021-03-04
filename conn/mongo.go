package conn

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

// initial data
func init() {
	host := "admin:admin@117.53.44.15:27017"
	dbName := "digileaps"
	session, err := mgo.Dial(host)
	log.Println("[*] Connecting To Database ...")
	if err != nil {
		log.Println("[-] Database Error Session :", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
	log.Println("[+] Connected To Database ...")
}

// GetDB for set connection
func GetDB() *mgo.Database {
	return db
}
