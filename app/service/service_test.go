package service

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/models"
	mocks "github.com/FrancoLiberali/stori_challenge/mocks/app/adapters"
)

func TestCalculateTotalBalance(t *testing.T) {
	service := Service{}
	tests := []struct {
		name string
		got  []models.Transaction
		want decimal.Decimal
	}{
		{"0 transactions returns 0", []models.Transaction{}, decimal.NewFromInt(0)},
		{"1 transaction returns first one", []models.Transaction{{Amount: decimal.NewFromFloat(60.5)}}, decimal.NewFromFloat(60.5)},
		{"multiple transaction returns sum", []models.Transaction{
			{Amount: decimal.NewFromFloat(60.5)},
			{Amount: decimal.NewFromFloat(-10.3)},
			{Amount: decimal.NewFromFloat(-20.46)},
			{Amount: decimal.NewFromFloat(10)},
		}, decimal.NewFromFloat(39.74)},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, service.CalculateTotalBalance(tt.got))
		})
	}
}

func TestCalculateTransactionsPerMonth(t *testing.T) {
	service := Service{}
	tests := []struct {
		name string
		got  []models.Transaction
		want []TransactionsPerMonth
	}{
		{"0 transactions returns empty", []models.Transaction{}, []TransactionsPerMonth{}},
		{"1 transaction returns 1 for one month", []models.Transaction{
			{Date: time.Date(time.Now().Year(), 7, 10, 0, 0, 0, 0, time.UTC)},
		}, []TransactionsPerMonth{
			{
				Month:  time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC),
				Amount: 1,
			},
		}},
		{"2 transactions in different months", []models.Transaction{
			{Date: time.Date(time.Now().Year(), 7, 10, 0, 0, 0, 0, time.UTC)},
			{Date: time.Date(time.Now().Year(), 8, 16, 0, 0, 0, 0, time.UTC)},
		}, []TransactionsPerMonth{
			{
				Month:  time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC),
				Amount: 1,
			},
			{
				Month:  time.Date(time.Now().Year(), 8, 1, 0, 0, 0, 0, time.UTC),
				Amount: 1,
			},
		}},
		{"multiples transactions in different months", []models.Transaction{
			{Date: time.Date(time.Now().Year(), 7, 15, 0, 0, 0, 0, time.UTC)},
			{Date: time.Date(time.Now().Year(), 7, 28, 0, 0, 0, 0, time.UTC)},
			{Date: time.Date(time.Now().Year(), 8, 2, 0, 0, 0, 0, time.UTC)},
			{Date: time.Date(time.Now().Year(), 8, 13, 0, 0, 0, 0, time.UTC)},
		}, []TransactionsPerMonth{
			{
				Month:  time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC),
				Amount: 2,
			},
			{
				Month:  time.Date(time.Now().Year(), 8, 1, 0, 0, 0, 0, time.UTC),
				Amount: 2,
			},
		}},
		{"multiples transactions in same month of different years", []models.Transaction{
			{Date: time.Date(time.Now().Year(), 7, 15, 0, 0, 0, 0, time.UTC)},
			{Date: time.Date(time.Now().Year(), 7, 28, 0, 0, 0, 0, time.UTC)},
			{Date: time.Date(2023, 7, 28, 0, 0, 0, 0, time.UTC)},
		}, []TransactionsPerMonth{
			{
				Month:  time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC),
				Amount: 1,
			},
			{
				Month:  time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC),
				Amount: 2,
			},
		}},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, service.CalculateTransactionsPerMonth(tt.got))
		})
	}
}

func TestCalculateAverageDebitAndCredit(t *testing.T) {
	service := Service{}
	tests := []struct {
		name   string
		got    []models.Transaction
		debit  string
		credit string
	}{
		{"0 transactions returns 0, 0", []models.Transaction{}, "0", "0"},
		{"1 debit returns that debit and 0", []models.Transaction{
			{Amount: decimal.NewFromFloat(-10.3)},
		}, "-10.3", "0"},
		{"1 credit returns that 0 and that credit", []models.Transaction{
			{Amount: decimal.NewFromFloat(60.5)},
		}, "0", "60.5"},
		{"1 debit and 1 credit returns that values", []models.Transaction{
			{Amount: decimal.NewFromFloat(60.5)},
			{Amount: decimal.NewFromFloat(-10.3)},
		}, "-10.3", "60.5"},
		{"multiple credits and debits", []models.Transaction{
			{Amount: decimal.NewFromFloat(60.5)},
			{Amount: decimal.NewFromFloat(-10.3)},
			{Amount: decimal.NewFromFloat(-20.46)},
			{Amount: decimal.NewFromFloat(10)},
		}, "-15.38", "35.25"},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			debit, credit := service.CalculateAverageDebitAndCredit(tt.got)
			assert.Equal(t, tt.debit, debit.String())
			assert.Equal(t, tt.credit, credit.String())
		})
	}
}

func TestProcessReturnsErrorIfCSVCantBeRead(t *testing.T) {
	mockCSVReader := mocks.NewCSVReader(t)
	service := Service{TransactionsReader: TransactionsReader{
		LocalCSVReader: mockCSVReader,
	}}

	mockCSVReader.On("Read", "not_found.csv").Return(nil, adapters.ErrReadingFile)

	err := service.Process("not_found.csv", "client@mail.com")
	require.ErrorIs(t, err, ErrReadingTransactions)
	require.ErrorContains(t, err, "error while reading file")
}

func TestProcessReturnsErrorIfCSVCantBeParsed(t *testing.T) {
	mockCSVReader := mocks.NewCSVReader(t)
	service := Service{TransactionsReader: TransactionsReader{
		LocalCSVReader: mockCSVReader,
	}}

	mockCSVReader.On("Read", "found.csv").Return([][]string{{"asd"}}, nil)

	err := service.Process("found.csv", "client@mail.com")
	require.ErrorIs(t, err, ErrReadingTransactions)
}

func TestReturnsErrorIfErrorIsProducedWhileSendingEmail(t *testing.T) {
	mockCSVReader := mocks.NewCSVReader(t)
	mockEmailSender := mocks.NewEmailSender(t)
	service := Service{
		TransactionsReader: TransactionsReader{
			LocalCSVReader: mockCSVReader,
		},
		EmailSender: mockEmailSender,
	}

	mockCSVReader.On("Read", "correct.csv").Return([][]string{{"0", "7/15", "+60.5"}}, nil)
	mockEmailSender.On("Send", "client@mail.com", emailSubject, `
Total balance is: 60.5
Number of transactions in July: 1
Average debit amount: 0
Average credit amount: 60.5
`).Return(adapters.ErrSendingEmail)

	err := service.Process("correct.csv", "client@mail.com")
	require.ErrorIs(t, err, adapters.ErrSendingEmail)
}

func TestProcessReturnsNilIfAllCorrect(t *testing.T) {
	mockCSVReader := mocks.NewCSVReader(t)
	mockEmailSender := mocks.NewEmailSender(t)
	service := Service{
		TransactionsReader: TransactionsReader{
			LocalCSVReader: mockCSVReader,
		},
		EmailSender: mockEmailSender,
	}

	mockCSVReader.On("Read", "correct.csv").Return([][]string{{"0", "7/15", "+60.5"}}, nil)
	mockEmailSender.On("Send", "client@mail.com", emailSubject, `
Total balance is: 60.5
Number of transactions in July: 1
Average debit amount: 0
Average credit amount: 60.5
`).Return(nil)

	err := service.Process("correct.csv", "client@mail.com")
	require.NoError(t, err)
}

func TestTransactionsPerMonthToString(t *testing.T) {
	tests := []struct {
		name string
		got  []TransactionsPerMonth
		want string
	}{
		{"empty list", []TransactionsPerMonth{}, ""},
		{"list with 1 transaction in one month", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 1},
		}, `Number of transactions in July: 1
`},
		{"list with multiple transaction in one month", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July: 2
`},
		{"list with multiple months", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			{Month: time.Date(time.Now().Year(), 8, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July: 2
Number of transactions in August: 2
`},
		{"list with multiple months of another year", []TransactionsPerMonth{
			{Month: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			{Month: time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July 2023: 2
Number of transactions in August 2023: 2
`},
		{"list with multiple years", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			{Month: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July: 2
Number of transactions in July 2023: 2
`},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, transactionsPerMonthToString(tt.got))
		})
	}
}