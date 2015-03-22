package main

import (
	"fmt"
	"log"
	"os"

	"github.com/natebrennand/twiliogo"
	"github.com/natebrennand/twiliogo/sms"
	"github.com/sendgrid/sendgrid-go"
)

func envVar(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("'%s' must be set in the environment", key)
	}
	return val
}

var (
	sendgridUser = envVar("SENDGRID_USER")
	sendgridPass = envVar("SENDGRID_PASS")
	recipient    = envVar("RECIPIENT_EMAIL")
	sender       = recipient
	sg           = sendgrid.NewSendGridClient(sendgridUser, sendgridPass)

	twilioFrom = envVar("TWILIO_NUMBER")
	twilioTo   = envVar("RECIPIENT_NUMBER")
	twilio     = twiliogo.NewAccountFromEnv()
)

func alert(url string) {
	// send an email alert
	message := sendgrid.NewMail()
	message.AddTo(recipient)
	message.SetFrom(recipient)
	message.SetSubject("BACCHANAL TICKET")
	message.SetText(fmt.Sprintf("Buy your ticket here: %s", url))

	if err := sg.Send(message); err != nil {
		log.Printf("ERROR (SENDGRID) => %s", err.Error())
	}
	log.Printf("Email sent to %s from %s", recipient, sender)

	// send a text notification
	resp, err := twilio.Sms.Send(sms.Post{
		From: twilioFrom,
		To:   twilioTo,
		Body: fmt.Sprintf("Buy your ticket here: %s", url),
	})
	if err != nil {
		log.Printf("SMS failed to send => {%s}", err)
	} else {
		log.Printf("SMS sent, %s", resp.Sid)
	}
}
