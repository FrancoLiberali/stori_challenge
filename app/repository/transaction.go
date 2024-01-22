package repository

import (
	"gorm.io/gorm"

	"github.com/FrancoLiberali/stori_challenge/app/models"
)

type ITransactionRepository interface {
	Create(db *gorm.DB, transactions []models.Transaction) error
}

type TransactionRepository struct{}

func (repository TransactionRepository) Create(db *gorm.DB, transactions []models.Transaction) error {
	return db.Create(transactions).Error
}
