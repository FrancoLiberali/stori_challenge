package service

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"github.com/shopspring/decimal"

	"github.com/FrancoLiberali/stori_challenge/app/adapters"
	"github.com/FrancoLiberali/stori_challenge/app/models"
)

type IEmailService interface {
	// Send formats the data received by parameters into the Stori mail and sends it by email
	Send(
		user *models.User,
		transactionBalance decimal.Decimal,
		transactionsPerMonth []TransactionsPerMonth,
		avgDebit, avgCredit decimal.Decimal,
	) error
}

type EmailService struct {
	EmailSender adapters.EmailSender
	Template    *template.Template
}

const emailSubject = "Stori transaction summary"

type transactionsPerMonthEmailData struct {
	Month string
	Value string
	IsOdd bool
}

type emailData struct {
	Date                     string
	UserBalance              string
	TransactionsBalance      string
	AvgDebit                 string
	AvgCredit                string
	TransactionsPerMonthList []transactionsPerMonthEmailData
}

// send formats the data received by parameters into the Stori mail and sends it by email
func (emailService EmailService) Send(
	user *models.User,
	transactionsBalance decimal.Decimal,
	transactionsPerMonth []TransactionsPerMonth,
	avgDebit, avgCredit decimal.Decimal,
) error {
	var htmlBuffer bytes.Buffer

	err := emailService.Template.Execute(&htmlBuffer, emailData{
		Date:                     time.Now().Format(time.DateTime),
		UserBalance:              user.Balance.String(),
		TransactionsBalance:      transactionsBalance.String(),
		AvgDebit:                 avgDebit.String(),
		AvgCredit:                avgCredit.String(),
		TransactionsPerMonthList: transactionsPerMonthToEmailData(transactionsPerMonth),
	})
	if err != nil {
		return fmt.Errorf("%w: %s", adapters.ErrSendingEmail, err.Error())
	}

	return emailService.EmailSender.Send(user.Email, emailSubject, htmlBuffer.String())
}

// transactionsPerMonthToEmailData transforms a list of transactions per month to
// a list of transactionsPerMonthEmailData that can be render into the email template
func transactionsPerMonthToEmailData(transactionsPerMonthList []TransactionsPerMonth) []transactionsPerMonthEmailData {
	currentYear := time.Now().Year()

	result := make([]transactionsPerMonthEmailData, 0, len(transactionsPerMonthList))

	for i, transactionsPerMonth := range transactionsPerMonthList {
		var monthText string

		// add year if not current year
		if transactionsPerMonth.Month.Year() != currentYear {
			monthText = fmt.Sprintf("%s %d", transactionsPerMonth.Month.Month().String(), transactionsPerMonth.Month.Year())
		} else {
			monthText = transactionsPerMonth.Month.Month().String()
		}

		result = append(result, transactionsPerMonthEmailData{
			Month: monthText,
			Value: strconv.Itoa(transactionsPerMonth.Amount),
			IsOdd: (i%2 == 0),
		})
	}

	return result
}
