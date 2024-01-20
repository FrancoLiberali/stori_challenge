package testintegration

import (
	"testing"

	"github.com/FrancoLiberali/stori_challenge/adapters"
	mocks "github.com/FrancoLiberali/stori_challenge/mocks/adapters"
	"github.com/FrancoLiberali/stori_challenge/service"
	"github.com/stretchr/testify/require"
)

func TestProcessLocalCSVFileSendsEmail(t *testing.T) {
	mockEmailSender := mocks.NewEmailSender(t)
	storiService := service.Service{
		CSVReader:   adapters.LocalCsvReader{},
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
