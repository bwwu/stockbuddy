package smtp

import (
	"fmt"
	"log"
	"net/smtp"
	"regexp"
	"strings"
)

var emailRE = regexp.MustCompile("\\w+@\\w+\\.\\w+")

type Email struct {
	Body, Subject string
	Recipients    []string
}

func (e *Email) Send(password string) error {
	auth := smtp.PlainAuth("", defaultSender.email, password, smtpServer)
	if err := smtp.SendMail(smtpServer+":587", auth, defaultSender.email, e.Recipients, getEmailContent(e)); err != nil {
		log.Printf("smtp: error '%v'", err)
		return err
	}
	log.Printf("smtp: email(s) sent to %s", strings.Join(e.Recipients, ","))
	return nil
}

var (
	smtpServer    = "smtp.gmail.com"
	defaultSender = sender{"tradingbot88@gmail.com", "Trade Bot"}
)

type sender struct {
	email, name string
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

// Given a comma-separated-list of emails given by a flag value, return a list of validated email
// addresses.
func ParseEmailsFromList(raw string) ([]string, error) {
	errPrefix := "smtp::ParseEmailFromList():"
	result := strings.Split(raw, ",")

	if len(result) == 0 {
		return nil, fmt.Errorf("%s empty email list", errPrefix)
	}

	// Validate email addresses.
	for _, email := range result {
		if !emailRE.MatchString(email) {
			return nil, fmt.Errorf(`%s invalid email "%s"`, errPrefix, email)
		}
	}
	return result, nil
}
