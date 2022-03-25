package main

import (
	"atm-api/database"
	"atm-api/handlers"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var server = handlers.Server{}

func main() {
	server.Init()

	err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	server.InitializeRoutes()

	log.Infoln("Listening on network " + server.NetworkAddress)
	log.Fatal(http.ListenAndServe(server.NetworkAddress, server.Router))
}
