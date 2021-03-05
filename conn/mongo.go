package conn

import (
	"flag"
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

// initial data
func init() {
	// host := "localhost:27017"
	// dbName := "example"
	host := flag.String("host", "localhost", "default host localhost")
	dbName := flag.String("coll", "example", "default DB example")

	flag.Parse()

	session, err := mgo.Dial(*host)
	log.Println("[*] Connecting To Database ...")
	if err != nil {
		log.Println("[-] Database Error Session :", err)
		os.Exit(2)
	}
	db = session.DB(*dbName)
	log.Println("[+] Connected To Database ...")
}

// GetDB for set connection
func GetDB() *mgo.Database {
	return db
}
