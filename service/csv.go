package service

import (
	"errors"
	"fmt"

	"github.com/FrancoLiberali/stori_challenge/models"
)

const (
	IDIndex                = 0
	IDHeader               = "Id"
	DateIndex              = 1
	DateHeader             = "Date"
	AmountIndex            = 2
	AmountHeader           = "Transaction"
	numberOfElementsPerRow = 3
)

var ErrParsingCsv = errors.New("error parsing transactions csv")

// csvRowsToTransactions transforms a list of csv rows into a list of models.Transaction,
// converting each cell of the row in the correct type.
//
// Returns the list of transactions or ErrParsingCsv if there is an error during type conversion
func csvRowsToTransactions(csvRows [][]string) ([]models.Transaction, error) {
	// create list with the same size as csvRows
	transactions := make([]models.Transaction, 0, len(csvRows))

	// transform each row into a models.Transaction
	for lineNumber, row := range csvRows {
		if len(row) != numberOfElementsPerRow {
			return nil, fmt.Errorf("%w: error parsing line %d: %d elements expected, got %d", ErrParsingCsv, lineNumber+1, numberOfElementsPerRow, len(row))
		}

		idString := row[IDIndex]

		id, err := parseTransactionID(idString)
		if err != nil {
			return nil, errorParsingCSV(IDHeader, idString, lineNumber)
		}

		dateString := row[DateIndex]

		date, err := parseDate(dateString)
		if err != nil {
			return nil, errorParsingCSV(DateHeader, dateString, lineNumber)
		}

		amountString := row[AmountIndex]

		amount, err := parseAmount(amountString)
		if err != nil {
			return nil, errorParsingCSV(AmountHeader, amountString, lineNumber)
		}

		transactions = append(transactions, models.Transaction{
			ID:     id,
			Date:   date,
			Amount: amount,
		})
	}

	return transactions, nil
}

// Returns a ErrParsingCsv with more information for the user
func errorParsingCSV(header string, stringValue string, lineNumber int) error {
	return fmt.Errorf(
		"%w: error parsing %s %q in line %d",
		ErrParsingCsv, header, stringValue, lineNumber+1,
	)
}
