package mailing

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	cfg "owlbytech/internal/config"
	dto "owlbytech/internal/dto/mailing"

	"gopkg.in/gomail.v2"
)

func SendPassword(send dto.PasswordMailingReq) error {

	html, err := parseHTML(send, "template-password")
	if err != nil {
		return fmt.Errorf("could not read HTML: %w", err)
	}

	err = SendEmail(send.Email, send.Subject, html)
	if err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}
func SendResetPassword(send *dto.ResetPasswordMailingReq) error {

	html, err := parseHTML(send, "template-reset-password")
	if err != nil {
		return fmt.Errorf("could not read HTML: %w", err)
	}

	err = SendEmail(send.Email, send.Subject, html)
	if err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}

func parseHTML(data interface{}, templatePath string) (string, error) {
	htmlContent, err := os.ReadFile(fmt.Sprintf("./internal/service/mailing/templates/%s.html", templatePath))
	if err != nil {
		return "", fmt.Errorf("error reading HMTL file: %w", err)
	}

	htmlTemplate := string(htmlContent)

	dataMap := make(map[string]interface{})
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling data: %w", err)
	}

	err = json.Unmarshal(dataBytes, &dataMap)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling data: %w", err)
	}

	for key, value := range dataMap {
		placeholder := fmt.Sprintf("{{%s}}", key)
		htmlTemplate = strings.ReplaceAll(htmlTemplate, placeholder, fmt.Sprintf("%v", value))
	}

	return htmlTemplate, nil
}

func SendEmail(email string, subject string, html string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.Env.EmailFrom)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", html)

	d := gomail.NewDialer(cfg.Env.EmailSMTPHost, cfg.Env.EmailSMTPPort, cfg.Env.EmailFrom, cfg.Env.EmailSMTPPassword)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %w", err)
	}

	return nil
}
