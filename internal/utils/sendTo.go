package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
	"time"

	"github.com/quangdvn/go-ec/global"
	"github.com/quangdvn/go-ec/pkg/settings"
	"go.uber.org/zap"
)

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

// var (
// 	MailServer   = global.Config.MailServer
// 	SMTPServer   = MailServer.Host
// 	SMTPPort     = MailServer.Port
// 	SMTPUsername = MailServer.Username
// 	SMTPPassword = MailServer.Password
// )

func getMailConfig() settings.MailServerSetting {
	return global.Config.MailServer
}

// TODO: Refactor to interface and use Sendgrid to send mail
func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, "; "))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

func SendTextEmail(to []string, from string, otp string) error {
	mailConfig := getMailConfig()
	emailContent := Mail{
		From:    EmailAddress{Name: "Test OTP", Address: from},
		To:      to,
		Subject: "OTP Verfication",
		Body:    fmt.Sprintf("Your OTP is %s.", otp),
	}
	messageEmail := BuildMessage(emailContent)

	// Send via SMTP
	auth := smtp.PlainAuth("", mailConfig.Username, mailConfig.Password, mailConfig.Host)
	start := time.Now()
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", mailConfig.Host, mailConfig.Port),
		auth,
		from,
		to,
		[]byte(messageEmail),
	)
	fmt.Println("Email sending took:", time.Since(start))
	if err != nil {
		global.Logger.Error("Email sent failed", zap.Error(err))
		return err
	}
	return nil
}

func SendOTPTemplateEmail(to []string, from string,
	templateName string,
	dataTemplate map[string]interface{}) error {
	htmlBody, err := getMailTemplate(templateName, dataTemplate)
	if err != nil {
		return err
	}
	return send(to, from, htmlBody)
}

func getMailTemplate(name string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)

	t := template.Must(template.New(name).ParseFiles(fmt.Sprintf("mailTemplates/%s", name)))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil

}

func send(to []string, from string, htmlTemplate string) error {
	mailConfig := getMailConfig()

	emailContent := Mail{
		From:    EmailAddress{Name: "Test OTP", Address: from},
		To:      to,
		Subject: "OTP Verfication",
		Body:    htmlTemplate,
	}
	messageEmail := BuildMessage(emailContent)

	// Send via SMTP
	auth := smtp.PlainAuth("", mailConfig.Username, mailConfig.Password, mailConfig.Host)
	start := time.Now()
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", mailConfig.Host, mailConfig.Port),
		auth,
		from,
		to,
		[]byte(messageEmail),
	)
	fmt.Println("Email sending took:", time.Since(start))
	if err != nil {
		global.Logger.Error("Email sent failed", zap.Error(err))
		return err
	}
	return nil
}
