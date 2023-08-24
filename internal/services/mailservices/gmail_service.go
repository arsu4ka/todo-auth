package mailservices

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

const (
	smtpAuthAddress = "smtp.gmail.com"
	smtpServerPort  = 587
)

type EmailSender struct {
	Message *gomail.Message
	dailer  *gomail.Dialer
	domain  string
}

// NewEmailSender
func NewEmailSender(domain, senderAddress, smtpPassword string) *EmailSender {
	m := gomail.NewMessage()
	dailer := gomail.NewDialer(smtpAuthAddress, smtpServerPort, senderAddress, smtpPassword)
	m.SetHeader("From", senderAddress)

	return &EmailSender{
		Message: m,
		dailer:  dailer,
		domain:  domain,
	}
}

// Send
func (es *EmailSender) SendVerificationLink(toAdress, toName, token string) error {
	var verificationLink string
	if es.domain[len(es.domain)-1] == '/' {
		verificationLink = es.domain + "auth/verify/" + token
	} else {
		verificationLink = es.domain + "/auth/verify/" + token
	}

	es.Message.SetHeader("To", toAdress)
	es.Message.SetHeader("Subject", "Todo App Email Verification")
	bodyFormat := "Hello, %s\n\nVisit this link to verify your email address: %s"
	es.Message.SetBody("text/plain", fmt.Sprintf(bodyFormat, toName, verificationLink))

	if err := es.dailer.DialAndSend(es.Message); err != nil {
		return err
	}

	return nil
}
