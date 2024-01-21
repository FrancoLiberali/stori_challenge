package app

import (
	"errors"
	"os"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/service"
)

//go:generate mockery --all --keeptree

const (
	emailPublicAPIKeyEnvVar  = "EMAIL_PUBLIC_API_KEY"  //nolint:gosec // just the env var name
	emailPrivateAPIKeyEnvVar = "EMAIL_PRIVATE_API_KEY" //nolint:gosec // just the env var name
)

var ErrEmailAPIKeyNotConfigured = errors.New("email api key env variables not configured")

// NewService creates a new service.Service.
// It reads the email public and private keys to inject them in the email sender.
func NewService() (*service.Service, error) {
	emailPublicAPIKey := os.Getenv(emailPublicAPIKeyEnvVar)
	emailPrivateAPIKey := os.Getenv(emailPrivateAPIKeyEnvVar)

	if emailPublicAPIKey == "" || emailPrivateAPIKey == "" {
		return nil, ErrEmailAPIKeyNotConfigured
	}

	return &service.Service{
		TransactionsReader: service.TransactionsReader{
			LocalCSVReader: adapters.LocalCSVReader{},
			S3CSVReader:    adapters.S3CSVReader{},
		},
		EmailSender: adapters.MailJetEmailSender{
			PublicAPIKey:  emailPublicAPIKey,
			PrivateAPIKey: emailPrivateAPIKey,
		},
	}, nil
}
