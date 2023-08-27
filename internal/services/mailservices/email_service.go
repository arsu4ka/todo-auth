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

func (es *EmailService) sendEmail(toAddress, subject, body string) error {
	es.Message.SetHeader("To", toAddress)
	es.Message.SetHeader("Subject", subject)
	es.Message.SetBody("text/plain", body)

	if err := es.dailer.DialAndSend(es.Message); err != nil {
		return err
	}
	return nil
}

// Send
func (es *EmailService) SendVerificationLink(toAdress, toName, token string) error {
	var verificationLink string
	if es.domain[len(es.domain)-1] == '/' {
		verificationLink = es.domain + "api/auth/verify/" + token
	} else {
		verificationLink = es.domain + "/api/auth/verify/" + token
	}

	subject := "Todo App Email Verification"
	bodyFormat := "Hello, %s\n\nVisit this link to verify your email address: %s"
	body := fmt.Sprintf(bodyFormat, toName, verificationLink)

	return es.sendEmail(toAdress, subject, body)
}

func (es *EmailService) SendResetLink(toAddress, toName, token string) error {
	var resetLink string
	if es.domain[len(es.domain)-1] == '/' {
		resetLink = es.domain + "api/auth/reset/" + token
	} else {
		resetLink = es.domain + "/api/auth/reset/" + token
	}

	subject := "Todo App Password Reset"
	bodyFormat := "Hello, %s\n\nTo reset your password send POST request to %s with your new password"
	body := fmt.Sprintf(bodyFormat, toName, resetLink)

	return es.sendEmail(toAddress, subject, body)
}
