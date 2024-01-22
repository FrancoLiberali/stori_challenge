package models

import (
	"github.com/shopspring/decimal"

	"github.com/FrancoLiberali/cql/model"
)

type User struct {
	model.UIntModel

	Email   string `gorm:"unique"`
	Balance decimal.Decimal
}
