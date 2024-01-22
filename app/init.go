package app

import (
	"embed"
	"errors"
	"html/template"
	"os"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/service"
)

//go:generate mockery --all --keeptree

const (
	emailPublicAPIKeyEnvVar  = "EMAIL_PUBLIC_API_KEY"  //nolint:gosec // just the env var name
	emailPrivateAPIKeyEnvVar = "EMAIL_PRIVATE_API_KEY" //nolint:gosec // just the env var name
	emailTemplate            = "html/email.html"
)

var ErrEmailAPIKeyNotConfigured = errors.New("email api key env variables not configured")

//go:embed html
var htmlFS embed.FS

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
		EmailService: service.EmailService{
			Template: template.Must(template.ParseFS(htmlFS, emailTemplate)),
			EmailSender: adapters.MailJetEmailSender{
				PublicAPIKey:  emailPublicAPIKey,
				PrivateAPIKey: emailPrivateAPIKey,
			},
		},
	}, nil
}
