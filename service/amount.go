package service

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

var ErrParsingAmount = errors.New("can't convert transaction amount to decimal")

func parseAmount(s string) (decimal.Decimal, error) {
	ans, err := decimal.NewFromString(s)
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("%w: %s", ErrParsingAmount, err.Error())
	}

	return ans, nil
}
