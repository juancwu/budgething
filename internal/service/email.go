package service

import (
	"context"
	"log/slog"

	"github.com/resend/resend-go/v2"
)

type EmailParams struct {
	From    string
	To      []string
	Subject string
	Text    string
}

type EmailClient interface {
	SendWithContext(ctx context.Context, params EmailParams) (string, error)
}

type ResendClient struct {
	client *resend.Client
}

func NewResendClient(apiKey string) *ResendClient {
	var client *resend.Client
	if apiKey != "" {
		client = resend.NewClient(apiKey)
	} else {
		slog.Warn("cannot initialize Resend client with empty api key")
		return nil
	}
	return &ResendClient{client: client}
}

func (c *ResendClient) SendWithContext(ctx context.Context, params EmailParams) (string, error) {
	res, err := c.client.Emails.SendWithContext(ctx, &resend.SendEmailRequest{
		From:    params.From,
		To:      params.To,
		Subject: params.Subject,
		Text:    params.Text,
	})
	if err != nil {
		return "", err
	}
	return res.Id, nil
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
