package main

import (
	"fmt"
	"log"
	"strings"
	"net/smtp"
)

var SMTP_SERVER = "smtp.gmail.com"

type Sender struct {
	email, name, password string
}

type Email struct {
	body, subject string
	recipients []string
	sender Sender
}

func (e *Email) Message() []byte {
	content := fmt.Sprintf("From:%s\nTo:%s\nSubject:%s\nContent-Type:text/html;charset=\"utf-8\"\n<html>%s</html>", e.sender.name, strings.Join(e.recipients, ","), e.subject, e.body)
	return []byte(content)
}

func main() {
	to := "brandonwu23@gmail.com"
	sender := Sender{"tradingbot88@gmail.com", "Trade Bot", "makemoney888$"}
	email := Email{"<b>Hi there!", "Hi!", []string{to}, sender}
	auth := smtp.PlainAuth(sender.name, sender.email, sender.password, SMTP_SERVER)
	if err := smtp.SendMail(SMTP_SERVER + ":587", auth, sender.email, email.recipients, email.Message()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	log.Println("Sent!")
}

