package mails

import (
	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/internal/utils"
)

// MailSender is the interface for sending email via any email provider (SMTP, SendGrid, etc).
// It allows flexibility and makes testing easier by abstracting the sending logic.
type MailSender interface {
	SendEmail(to []string, from, subject, htmlBody string) error
}

func SendTextEmail(to []string, from string, otp string) error {
	var mailer MailSender = NewGmailMailer() // or swap in SendGridMailer
	return mailer.SendEmail(to, from, "OTP Verfication", otp)
}

func SendOTPTemplateEmail(to []string, from string,
	templateName string,
	dataTemplate map[string]interface{}) error {
	htmlBody, err := utils.GetMailTemplate(templateName, dataTemplate)
	if err != nil {
		return err
	}
	var mailer MailSender
	switch global.Config.MailServer.Provider {
	case "sendgrid":
		mailer = NewSendGridMailer()
	default:
		mailer = NewGmailMailer()
	}
	return mailer.SendEmail(to, from, "OTP Verfication", htmlBody)
}
