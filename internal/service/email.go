package service

import "context"

type EmailParams struct {
	From    string
	To      []string
	Subject string
	Text    string
}

type EmailClient interface {
	SendWithContext(ctx context.Context, params EmailParams) (string, error)
}

type EmailService struct {
	client    EmailClient
	fromEmail string
	isDev     bool
	appURL    string
	appName   string
}

func NewEmailService(client EmailClient, fromEmail, appURL, appName string, isDev bool) *EmailService {
	return &EmailService{
		client:    client,
		fromEmail: fromEmail,
		isDev:     isDev,
		appURL:    appURL,
		appName:   appName,
	}
}
