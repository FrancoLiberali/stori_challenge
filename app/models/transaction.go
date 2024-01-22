package models

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/FrancoLiberali/cql/model"
)

const AmountBytes = 64

type Transaction struct {
	model.UIntModel

	IDInFile uint
	FileName string
	Date     time.Time
	Amount   decimal.Decimal
}
