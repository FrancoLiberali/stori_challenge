package service

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
)

type EmailService struct {
	EmailSender adapters.EmailSender
}

const (
	emailMessage = `
Total balance is: %s
%sAverage debit amount: %s
Average credit amount: %s
`
	emailSubject = "Stori transaction summary"
)

func (emailService EmailService) send(
	destinationEmail string,
	totalBalance decimal.Decimal,
	transactionsPerMonth []TransactionsPerMonth,
	avgDebit, avgCredit decimal.Decimal,
) error {
	return emailService.EmailSender.Send(destinationEmail, emailSubject, fmt.Sprintf(
		emailMessage,
		totalBalance,
		transactionsPerMonthToString(transactionsPerMonth),
		avgDebit,
		avgCredit,
	))
}

// transactionsPerMonthToString transforms a list of transactions per month to string that can be sent to the user by email
func transactionsPerMonthToString(transactionsPerMonth []TransactionsPerMonth) string {
	result := ""

	for _, transactionPerMonth := range transactionsPerMonth {
		// add year if not current year
		if transactionPerMonth.Month.Year() != time.Now().Year() {
			result += fmt.Sprintf(
				"Number of transactions in %s %d: %d\n",
				transactionPerMonth.Month.Month().String(),
				transactionPerMonth.Month.Year(),
				transactionPerMonth.Amount,
			)
		} else {
			result += fmt.Sprintf(
				"Number of transactions in %s: %d\n",
				transactionPerMonth.Month.Month().String(),
				transactionPerMonth.Amount,
			)
		}
	}

	return result
}
