package main

import (
	"log"
	"time"

	"github.com/sanyog2491/bookings2/internal/models"

	mail "github.com/xhit/go-simple-mail/v2"
)

func listetnformail() {
	//function that executes in the background
	go func() {
		for {
			msg := <-app.Mailchan
			sendMsg(msg)
		}
	}()
}

// this just sends an email
func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}
	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("email sent")
	}
}
