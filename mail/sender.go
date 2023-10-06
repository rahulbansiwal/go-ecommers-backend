package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type GmailSender struct {
	name              string
	fromEmailAddr     string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddr string, fromEmailPassword string) *GmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddr:     fromEmailAddr,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s %s", sender.name, sender.fromEmailAddr)
	e.Subject = subject
	e.Cc = cc
	e.Bcc = bcc
	e.HTML = []byte(content)
	e.To = to
	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("error while attaching file")
		}
	}
	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddr, sender.fromEmailPassword, smtpAuthAddress)
	err := e.Send(smtpServerAddress, smtpAuth)
	if err != nil {
		return err
	}
	return nil
}


// Add Redis Asynq to send verify email
