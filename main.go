package main

import (
	// Import the pq driver so that it can register itself with the database/sql
	// package. Note that we alias this import to the blank identifier, to stop the Go
	// compiler complaining that the package isn't being used.
	"log"

	config "github.com/dhxmo/shop-stop-go/config"
	_ "github.com/lib/pq"
)

func main() {
	_, err := config.ConnectDB()
	if err != nil {
		log.Fatal("error in db connection", err)
	}
	log.Println("successful db connection")
}
