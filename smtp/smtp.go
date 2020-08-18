package smtp

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type Email struct {
	Body, Subject string
	Recipients    []string
}

func (e *Email) Send() error {
	auth := smtp.PlainAuth(defaultSender.name, defaultSender.email, defaultSender.password, smtpServer)
	if err := smtp.SendMail(smtpServer+":587", auth, defaultSender.email, e.Recipients, getEmailContent(e)); err != nil {
		log.Print(err)
		return err
	}
	log.Printf("smtp: email(s) sent to %s", strings.Join(e.Recipients, ","))
	return nil
}

var (
	smtpServer    = "smtp.gmail.com"
	defaultSender = sender{"tradingbot88@gmail.com", "Trade Bot", "makemoney888$"}
)

type sender struct {
	email, name, password string
}

func getEmailContent(e *Email) []byte {
	content := fmt.Sprintf(
		"From:%s\nTo:%s\nSubject:%s\nContent-Type:text/html;charset=\"utf-8\"\n<html>\n%s\n</html>",
		defaultSender.name,
		strings.Join(e.Recipients, ","),
		e.Subject,
		e.Body,
	)
	return []byte(content)
}
