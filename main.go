package main

import (
	"net/http"
	"os"

	"github.com/jenkins-x/jx-cdevents-adapter/handlers"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func handleRequests() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	port, found := os.LookupEnv("ADAPTER_PORT")
	if !found || port == "" {
		port = "80"
	}

	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	handleRequests()
}
