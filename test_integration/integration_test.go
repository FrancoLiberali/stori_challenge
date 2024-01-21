package testintegration

import (
	"testing"

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
		EmailSender: mockEmailSender,
	}

	mockEmailSender.On("Send", "client@mail.com", "Stori transaction summary", `
Total balance is: 39.74
Number of transactions in July: 2
Number of transactions in August: 2
Average debit amount: -15.38
Average credit amount: 35.25
`).Return(nil)

	err := storiService.Process("../txns1.csv", "client@mail.com")
	require.NoError(t, err)
}

func TestProcessS3CSVFileSendsEmail(t *testing.T) {
	mockEmailSender := mocks.NewEmailSender(t)
	mockS3Reader := mocks.NewCSVReader(t)
	storiService := service.Service{
		TransactionsReader: service.TransactionsReader{
			S3CSVReader: mockS3Reader,
		},
		EmailSender: mockEmailSender,
	}

	mockS3Reader.On("Read", "fl-stori-challenge/txns1.csv").Return(
		[][]string{
			{"0", "7/15", "+60.5"},
			{"1", "7/28", "-10.3"},
			{"2", "8/2", "-20.46"},
			{"3", "8/13", "+10"},
		}, nil,
	)
	mockEmailSender.On("Send", "client@mail.com", "Stori transaction summary", `
Total balance is: 39.74
Number of transactions in July: 2
Number of transactions in August: 2
Average debit amount: -15.38
Average credit amount: 35.25
`).Return(nil)

	err := storiService.Process("s3://fl-stori-challenge/txns1.csv", "client@mail.com")
	require.NoError(t, err)
}
