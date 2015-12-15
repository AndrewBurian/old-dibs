package main

import (
	"fmt"
	"github.com/andrewburian/crypter"
	"net/http"
)

func main() {

	// static content
	http.Handle("/", http.FileServer(http.Dir("./views/")))

	// Setup crypto
	key := []byte("cryptokey0123456")
	mac := []byte("cryptomackey")
	iv := []byte("1234567890123456")

	fmt.Println("Setting up crypto")
	crypto, err := crypter.NewCrypter(key, mac, iv)
	if err != nil {
		panic(err)
	}

	// Setup database
	dbuser := "dibsagent"
	dbpassword := "agentpassword"
	dbname := "dibs"

	fmt.Println("Connecting to database")
	db, err := newDbManager(dbuser, dbpassword, dbname, 2)
	if err != nil {
		panic(err)
	}

	// REST endpoint handlers
	auth := &authHandler{crypto, db}
	http.Handle("/auth", auth)

	fmt.Println("Starting server!")
	http.ListenAndServe(":8080", nil)
}
