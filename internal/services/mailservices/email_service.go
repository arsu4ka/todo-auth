package mailservices

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

const (
	smtpAuthAddress = "smtp.office365.com"
	smtpServerPort  = 587
)

type EmailService struct {
	Message *gomail.Message
	dailer  *gomail.Dialer
	domain  string
}

// NewEmailSender
func NewEmailService(apiUrl, smtpEmail, smtpPassword string) *EmailService {
	m := gomail.NewMessage()
	dailer := gomail.NewDialer(smtpAuthAddress, smtpServerPort, smtpEmail, smtpPassword)
	dailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m.SetHeader("From", smtpEmail)

	return &EmailService{
		Message: m,
		dailer:  dailer,
		domain:  apiUrl,
	}
}

// Send
func (es *EmailService) SendVerificationLink(toAdress, toName, token string) error {
	var verificationLink string
	if es.domain[len(es.domain)-1] == '/' {
		verificationLink = es.domain + "api/auth/verify/" + token
	} else {
		verificationLink = es.domain + "/api/auth/verify/" + token
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
