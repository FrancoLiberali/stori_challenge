package service

import (
	"log"
	"sort"
	"time"

	"github.com/elliotchance/pie/v2"
	"github.com/shopspring/decimal"

	"github.com/FrancoLiberali/stori_challenge/adapters"
	"github.com/FrancoLiberali/stori_challenge/models"
)

type Service struct {
	CSVReader adapters.CSVReader
}

type TransactionsPerMonth struct {
	Month  time.Time
	Amount int
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
	transactionsPerMonth := service.CalculateTransactionsPerMonth(transactions)

	log.Println(totalBalance)
	log.Println(transactionsPerMonth)

	log.Println(destinationEmail)

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

// CalculateTransactionsPerMonth calculates the amount of transactions
// for each month present in the list of transactions.
// The returned list is in ascending order by month.
func (service Service) CalculateTransactionsPerMonth(transactions []models.Transaction) []TransactionsPerMonth {
	// group by year and month
	grouped := pie.GroupBy(transactions, func(transaction models.Transaction) time.Time {
		return time.Date(transaction.Date.Year(), transaction.Date.Month(), 1, 0, 0, 0, 0, time.UTC)
	})

	ans := make([]TransactionsPerMonth, 0, len(grouped))

	// count amount of transactions per group
	for month, transactions := range grouped {
		ans = append(ans, TransactionsPerMonth{Month: month, Amount: len(transactions)})
	}

	// order by month
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Month.Compare(ans[j].Month) == -1
	})

	return ans
}
