package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/models"
)

const (
	s3Prefix               = "s3://"
	IDIndex                = 0
	IDHeader               = "Id"
	DateIndex              = 1
	DateHeader             = "Date"
	AmountIndex            = 2
	AmountHeader           = "Transaction"
	numberOfElementsPerRow = 3
)

type TransactionsReader struct {
	LocalCSVReader adapters.CSVReader
	S3CSVReader    adapters.CSVReader
}

var ErrReadingTransactions = errors.New("error parsing transactions csv")

// Read reads a CSV file that contains a list of transactions.
//
// If the csvFileName starts with 's3://', it reads the csv from s3. Otherwise, it reads it from local folder.
//
// Returns the list of transactions of the CSV file
// or ErrReadingTransactions if an error is produced
func (reader TransactionsReader) Read(csvFileName string) ([]models.Transaction, error) {
	var csvReader adapters.CSVReader

	if strings.HasPrefix(csvFileName, s3Prefix) {
		csvReader = reader.S3CSVReader
		csvFileName = csvFileName[len(s3Prefix):]
	} else {
		csvReader = reader.LocalCSVReader
	}

	csvRows, err := csvReader.Read(csvFileName)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrReadingTransactions, err.Error())
	}

	return reader.parse(csvRows, csvFileName)
}

// Parse transforms a list of csv rows into a list of models.Transaction,
// converting each cell of the row in the correct type.
//
// Returns the list of transactions or ErrParsingCsv if there is an error during type conversion
func (reader TransactionsReader) parse(csvRows [][]string, csvFileName string) ([]models.Transaction, error) {
	// create list with the same size as csvRows
	transactions := make([]models.Transaction, 0, len(csvRows))

	// transform each row into a models.Transaction
	for lineNumber, row := range csvRows {
		if len(row) != numberOfElementsPerRow {
			return nil, fmt.Errorf("%w: error parsing line %d: %d elements expected, got %d", ErrReadingTransactions, lineNumber+1, numberOfElementsPerRow, len(row))
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
			IDInFile: id,
			FileName: csvFileName,
			Date:     date,
			Amount:   amount,
		})
	}

	return transactions, nil
}

// Returns a ErrParsingCsv with more information for the user
func errorParsingCSV(header string, stringValue string, lineNumber int) error {
	return fmt.Errorf(
		"%w: error parsing %s %q in line %d",
		ErrReadingTransactions, header, stringValue, lineNumber+1,
	)
}
