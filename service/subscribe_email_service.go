package service

import (
	"encoding/json"
	"fmt"
	"github.com/ac-arconomy/nbd-web-go/model"
	"net/http"
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

	// Do something with the Person struct...
	fmt.Fprintf(w, "Person: %+v", lead)


	//w.Write([]byte(`{"message": "hello world"}`))
}

