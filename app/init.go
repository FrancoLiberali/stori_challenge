package app

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/repository"
	"github.com/FrancoLiberali/stori_challenge/app/service"
)

//go:generate mockery --all --keeptree

const (
	EmailPublicAPIKeyEnvVar  = "EMAIL_PUBLIC_API_KEY"  //nolint:gosec // just the env var name
	EmailPrivateAPIKeyEnvVar = "EMAIL_PRIVATE_API_KEY" //nolint:gosec // just the env var name
	DBURLEnvVar              = "DB_URL"
	DBPortEnvVar             = "DB_PORT"
	DBUserEnvVar             = "DB_USER"
	DBPasswordEnvVar         = "DB_PASSWORD"
	DBNameEnvVar             = "DB_NAME"
	DBSSLEnvVar              = "DB_SSL"
	emailTemplate            = "html/email.html"
)

var (
	ErrEmailAPIKeyNotConfigured = errors.New("email api key env variables not configured")
	ErrDatabaseNotConfigured    = errors.New("database env variables not configured")
)

//go:embed html
var htmlFS embed.FS

// NewService creates a new service.Service.
// It reads the email public and private keys to inject them in the email sender.
func NewService(db *gorm.DB) (*service.Service, error) {
	emailPublicAPIKey := os.Getenv(EmailPublicAPIKeyEnvVar)
	emailPrivateAPIKey := os.Getenv(EmailPrivateAPIKeyEnvVar)

	if emailPublicAPIKey == "" || emailPrivateAPIKey == "" {
		return nil, ErrEmailAPIKeyNotConfigured
	}

	return &service.Service{
		TransactionsReader: service.TransactionsReader{
			LocalCSVReader: adapters.LocalCSVReader{},
			S3CSVReader:    adapters.S3CSVReader{},
		},
		TransactionService: service.TransactionService{
			DB:                    db,
			UserRepository:        repository.UserRepository{},
			TransactionRepository: repository.TransactionRepository{},
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

func NewDBConnection() (*gorm.DB, error) {
	dbURL := os.Getenv(DBURLEnvVar)
	dbPort := os.Getenv(DBPortEnvVar)
	dbUser := os.Getenv(DBUserEnvVar)
	dbPassword := os.Getenv(DBPasswordEnvVar)
	dbName := os.Getenv(DBNameEnvVar)
	dbSSL := os.Getenv(DBSSLEnvVar)

	if dbURL == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbSSL == "" {
		return nil, ErrDatabaseNotConfigured
	}

	return cql.Open(postgres.Open(
		fmt.Sprintf(
			"user=%s password=%s host=%s port=%s sslmode=%s dbname=%s",
			dbUser, dbPassword, dbURL, dbPort, dbSSL, dbName,
		),
	))
}
