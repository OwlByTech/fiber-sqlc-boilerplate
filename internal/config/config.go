package config

import (
	"fmt"
	"os"
	"strconv"
)

type EnvStruct struct {
	EmailFrom         string
	EmailSMTPHost     string
	EmailSMTPPort     int
	EmailSMTPPassword string
	JWTSecret         string
	JWTSecretPassword string
	WebUrl            string
}

var Env *EnvStruct

func LoadConfig() error {
	portStr := os.Getenv("EMAIL_SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP port %q: %w", portStr, err)
	}

	Env = &EnvStruct{
		EmailFrom:         os.Getenv("EMAIL_FROM"),
		EmailSMTPHost:     os.Getenv("EMAIL_SMTP_HOST"),
		EmailSMTPPort:     port,
		EmailSMTPPassword: os.Getenv("EMAIL_SMTP_PASSWORD"),
		JWTSecret:         os.Getenv("JWT_SECRET"),
		JWTSecretPassword: os.Getenv("JWT_SECRET_PASSWORD"),
		WebUrl:            os.Getenv("WEB_URL"),
	}

	return err
}
