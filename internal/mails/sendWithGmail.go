package mails

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/pkg/settings"
	"go.uber.org/zap"
)

// ============================
// Gmail SMTP Implementation
// ============================

type EmailAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

type GmailMailer struct {
	Config settings.GmailServerSetting
}

func NewGmailMailer() *GmailMailer {
	return &GmailMailer{Config: global.Config.Gmail}
}

func buildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, "; "))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func (gm *GmailMailer) SendEmail(to []string, from, subject, htmlBody string) error {
	emailContent := Mail{
		From:    EmailAddress{Name: "Test OTP", Address: from},
		To:      to,
		Subject: subject,
		Body:    htmlBody,
	}
	messageEmail := buildMessage(emailContent)

	auth := smtp.PlainAuth("", gm.Config.Username, gm.Config.Password, gm.Config.Host)
	start := time.Now()
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", gm.Config.Host, gm.Config.Port),
		auth,
		from,
		to,
		[]byte(messageEmail),
	)
	fmt.Println("Email sending took:", time.Since(start))
	if err != nil {
		global.Logger.Error("SMTP email send failed", zap.Error(err))
		return err
	}
	global.Logger.Info("Email sent via Gmail SMTP successfully")
	return nil
}
