package mails

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/quangdvn/go-ec/global"
	"github.com/sendgrid/sendgrid-go"
	"go.uber.org/zap"
)

// ============================
// SendGrid Implementation
// ============================

type SendGridMailer struct {
	ApiKey string
}

func NewSendGridMailer() *SendGridMailer {
	return &SendGridMailer{ApiKey: global.Config.SendGrid.ApiKey}
}

func (sg *SendGridMailer) SendEmail(to []string, from, subject, htmlBody string) error {
	fmt.Println("Sending email via SendGrid, ", sg.ApiKey)
	request := sendgrid.GetRequest(sg.ApiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"

	personalization := []map[string]interface{}{
		{
			"to":      convertToEmailObjects(to),
			"subject": subject,
		},
	}

	body := map[string]interface{}{
		"personalizations": personalization,
		"from": map[string]string{
			"email": from,
		},
		"content": []map[string]string{
			{
				"type":  "text/html",
				"value": htmlBody,
			},
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		global.Logger.Error("Failed to marshal SendGrid email body", zap.Error(err))
		return err
	}
	request.Body = bodyBytes
	start := time.Now()
	response, err := sendgrid.API(request)

	if err != nil {
		global.Logger.Error("SendGrid API call failed", zap.Error(err))
		return err
	}

	if response.StatusCode >= 400 {
		global.Logger.Error("SendGrid API returned error status", zap.Int("status", response.StatusCode))
		return fmt.Errorf("sendgrid failed with status code %d", response.StatusCode)
	}
	fmt.Println("Email sending took:", time.Since(start))
	global.Logger.Info("Email sent via SendGrid successfully")
	return nil
}

func convertToEmailObjects(addresses []string) []map[string]string {
	result := make([]map[string]string, len(addresses))
	for i, addr := range addresses {
		result[i] = map[string]string{"email": addr}
	}
	return result
}
