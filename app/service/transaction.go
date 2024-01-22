package service

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/FrancoLiberali/cql"
	"github.com/FrancoLiberali/stori_challenge/app/models"
	"github.com/FrancoLiberali/stori_challenge/app/repository"
)

type ITransactionService interface {
	// Apply creates the transactions received by parameter and applies its total balance to the user's balance
	Apply(email string, transactions []models.Transaction, transactionsBalance decimal.Decimal) (*models.User, error)
}

type TransactionService struct {
	DB                    *gorm.DB
	UserRepository        repository.IUserRepository
	TransactionRepository repository.ITransactionRepository
}

var ErrApplyingTransactions = errors.New("error applying transactions")

// Apply creates the transactions received by parameter and applies its total balance to the user's balance
func (service TransactionService) Apply(
	email string,
	transactions []models.Transaction,
	transactionsBalance decimal.Decimal,
) (*models.User, error) {
	return cql.Transaction(service.DB, func(tx *gorm.DB) (*models.User, error) {
		user, err := service.UserRepository.GetByEmail(tx, email)

		if errors.Is(err, cql.ErrObjectNotFound) {
			user = &models.User{Email: email, Balance: decimal.Zero}
		} else if err != nil {
			return nil, errApplyingTransactions(email, transactions, err)
		}

		user.Balance = user.Balance.Add(transactionsBalance)

		err = service.UserRepository.Save(tx, user)
		if err != nil {
			return nil, errApplyingTransactions(email, transactions, err)
		}

		err = service.TransactionRepository.Create(tx, transactions)
		if err != nil {
			return nil, errApplyingTransactions(email, transactions, err)
		}

		return user, nil
	})
}

func errApplyingTransactions(email string, transactions []models.Transaction, internalError error) error {
	return fmt.Errorf("%w %v to user %s: %s", ErrApplyingTransactions, transactions, email, internalError.Error())
}
