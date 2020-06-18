package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ac-arconomy/nbd-web-go/model"
	"log"
	"net/http"
	"os"
	"strings"
)

type SubscribeEmail struct{}

func (s *SubscribeEmail) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering SubscribeEmail..")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var lead model.Lead
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&lead)
	if err != nil {
		log.Println("Error SubscribeEmail.."  + err.Error())

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("SubscribeEmail Received Lead %+v\n", lead)

	client := &http.Client{}

	insightlyUrl := os.Getenv("INSIGHTLY_URL")
	insightlyApiKey := os.Getenv("INSIGHTLY_APIKEY")

	names := strings.Split(lead.FirstName, " ")
	lead.FirstName = names[0]

	var lastName = ""
	for i, part := range names {
		if i != 0 {
			lastName = lastName + " " + part
		} else  {}

	}
	lead.LastName = strings.TrimSpace(lastName)

	leadJson, err := json.Marshal(lead)
	req, err := http.NewRequest("POST", insightlyUrl, bytes.NewReader(leadJson))
	req.Header.Add("User-Agent", "nbd-http-client")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization","Basic " + basicAuth(insightlyApiKey,""))

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error SubscribeEmail.."  + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		defer resp.Body.Close()
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}


