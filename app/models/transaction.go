package models

import (
	"time"

	"github.com/shopspring/decimal"
)

const AmountBytes = 64

type Transaction struct {
	ID     uint
	Date   time.Time
	Amount decimal.Decimal
}
