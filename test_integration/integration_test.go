package testintegration

import (
	"html/template"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	mocks "github.com/FrancoLiberali/stori_challenge/app/mocks/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/service"
)

func TestProcessLocalCSVFileSendsEmail(t *testing.T) {
	mockEmailSender := mocks.NewEmailSender(t)
	storiService := service.Service{
		TransactionsReader: service.TransactionsReader{
			LocalCSVReader: adapters.LocalCSVReader{},
		},
		EmailService: service.EmailService{
			Template:    template.Must(template.ParseFiles("../app/html/email.html")),
			EmailSender: mockEmailSender,
		},
	}

	mockEmailSender.On(
		"Send",
		"client@mail.com",
		"Stori transaction summary",
		mock.MatchedBy(func(emailBody string) bool {
			return strings.Contains(emailBody, "39.74") && strings.Contains(emailBody, "-15.38") && strings.Contains(emailBody, "35.25")
		}),
	).Return(nil)

	err := storiService.Process("../data/txns1.csv", "client@mail.com")
	require.NoError(t, err)
}

func TestProcessS3CSVFileSendsEmail(t *testing.T) {
	mockEmailSender := mocks.NewEmailSender(t)
	mockS3Reader := mocks.NewCSVReader(t)
	storiService := service.Service{
		TransactionsReader: service.TransactionsReader{
			S3CSVReader: mockS3Reader,
		},
		EmailService: service.EmailService{
			Template:    template.Must(template.ParseFiles("../app/html/email.html")),
			EmailSender: mockEmailSender,
		},
	}

	mockS3Reader.On("Read", "fl-stori-challenge/txns1.csv").Return(
		[][]string{
			{"0", "7/15", "+60.5"},
			{"1", "7/28", "-10.3"},
			{"2", "8/2", "-20.46"},
			{"3", "8/13", "+10"},
		}, nil,
	)
	mockEmailSender.On(
		"Send",
		"client@mail.com",
		"Stori transaction summary",
		mock.MatchedBy(func(emailBody string) bool {
			return strings.Contains(emailBody, "39.74") && strings.Contains(emailBody, "-15.38") && strings.Contains(emailBody, "35.25")
		}),
	).Return(nil)

	err := storiService.Process("s3://fl-stori-challenge/txns1.csv", "client@mail.com")
	require.NoError(t, err)
}
