package service

import (
	"log"

	"github.com/FrancoLiberali/stori_challenge/adapters"
	"github.com/FrancoLiberali/stori_challenge/models"
	"github.com/elliotchance/pie/v2"
	"github.com/shopspring/decimal"
)

type Service struct {
	CSVReader adapters.CSVReader
}

// Process read csv file called csvFileName,
// calculates total balance, number of transactions grouped by month
// and average credit and debit
// and send this information by email to destinationEmail
func (service Service) Process(csvFileName, destinationEmail string) error {
	csvRows, err := service.CSVReader.Read(csvFileName)
	if err != nil {
		return err
	}

	transactions, err := csvRowsToTransactions(csvRows)
	if err != nil {
		return err
	}

	totalBalance := service.CalculateTotalBalance(transactions)

	log.Println(totalBalance)

	return nil
}

// CalculateTotalBalance calculates the total balance from a list of transactions as the sum of all transactions
func (service Service) CalculateTotalBalance(transactions []models.Transaction) decimal.Decimal {
	amounts := pie.Map(transactions, func(transaction models.Transaction) decimal.Decimal {
		return transaction.Amount
	})

	if len(amounts) == 0 {
		return decimal.NewFromInt(0)
	}

	if len(amounts) == 1 {
		return amounts[0]
	}

	return decimal.Sum(
		amounts[0],
		amounts[1:]...,
	)
}
