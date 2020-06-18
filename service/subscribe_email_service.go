package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/ac-arconomy/nbd-web-go/model"
	"net/http"
	"os"
)

type SubscribeEmail struct{}

func (s *SubscribeEmail) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var lead model.Lead
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&lead)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := &http.Client{}

	leadJson, err := json.Marshal(lead)

	insightlyUrl := os.Getenv("INSIGHTLY_URL")
	insightlyApiKey := os.Getenv("INSIGHTLY_APIKEY")

	req, err := http.NewRequest("POST", insightlyUrl, bytes.NewReader(leadJson))
	req.Header.Add("User-Agent", "nbd-http-client")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization","Basic " + basicAuth(insightlyApiKey,""))

	resp, err := client.Do(req)
	defer resp.Body.Close()
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}


