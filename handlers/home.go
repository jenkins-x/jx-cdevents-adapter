package handlers

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func home(w http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatalf("Could not read request body: %s", err)
	}

	log.Printf("Payload Received: %s", string(body))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
