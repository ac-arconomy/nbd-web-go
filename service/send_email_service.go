package service

import (
	"encoding/json"
	"fmt"
	"github.com/ac-arconomy/nbd-web-go/model"
	"log"
	"net/http"
	"net/smtp"
)

type SendEmail struct{}

func (s *SendEmail) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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


	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "ac@arconomy.digital", "1lovejaffacakes", "smtp.gmail.com")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"glenn@pringle.com.au"}
	msg := []byte("To: glenn@pringle.com.au\r\n" +
		"Subject: the subjeect?\r\n" +
		"\r\n" +
		"testing the mail\r\n")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "glenn@pringle.com.au", to, msg)
	if err != nil {
		log.Fatal(err)
	}


	//spring.mail.protocol=smtp
	//spring.mail.host=smtp.gmail.com
	//spring.mail.port=587
	//spring.mail.username=ac@arconomy.digital
	//spring.mail.password=1lovejaffacakes
	//spring.mail.properties.mail.smtp.auth = true
	//spring.mail.properties.mail.smtp.starttls.enable = tr
	//w.Write([]byte(`{"message": "hello world"}`))
}
