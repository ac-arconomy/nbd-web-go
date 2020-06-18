package main

import (
	"github.com/ac-arconomy/nbd-web-go/service"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	sendEmail := &service.SendEmail{}
	http.Handle("/send-email", sendEmail)

	subscribeEmail := &service.SubscribeEmail{}
	http.Handle("/subscribe-email", subscribeEmail)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}