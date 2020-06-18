package service

import (
	"encoding/json"
	"fmt"
	"github.com/ac-arconomy/nbd-web-go/model"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type SendEmail struct{}

func (s *SendEmail) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering SendEmail..")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var lead model.Lead
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&lead)
	if err != nil {
		log.Println("Error SendEmail.."  + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("SendEmail Received Lead %+v\n", lead)

	var smtpUsername = os.Getenv("SMTP_USERNAME")
	var smtpPassword = os.Getenv("SMTP_PASSWORD")
	var smtpHost = os.Getenv("SMTP_HOST")
	var smtpPort = os.Getenv("SMTP_PORT")

	// Choose auth method and set it up
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	// Here we do it all: connect to our server, set up a message and send it
	//ashlee@naturalbydesign.com.au
	var sendTo = os.Getenv("SMTP_SEND_TO")
	to := []string{sendTo}

	var subject = "Subject: " + "Contact form: " + lead.FirstName + " " + lead.LastName
	msg := []byte("To: "  + sendTo  + "\r\n" +
		subject  + "\r\n" +
		lead.Message + "\r\n")
	err = smtp.SendMail(smtpHost + ":" + smtpPort, auth, lead.Email, to, msg)
	if err != nil {
		log.Println("Error smtp SendMail() "  + err.Error())
		log.Fatal(err)
	}

}
