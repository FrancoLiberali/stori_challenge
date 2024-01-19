package service

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/FrancoLiberali/stori_challenge/models"
)

func TestCalculateTotalBalance(t *testing.T) {
	service := Service{}
	tests := []struct {
		name string
		got  []models.Transaction
		want decimal.Decimal
	}{
		{"0 transactions returns 0", []models.Transaction{}, decimal.NewFromInt(0)},
		{"1 transaction returns first one", []models.Transaction{{Amount: decimal.NewFromFloat(60.5)}}, decimal.NewFromFloat(60.5)},
		{"multiple transaction returns sum", []models.Transaction{
			{Amount: decimal.NewFromFloat(60.5)},
			{Amount: decimal.NewFromFloat(-10.3)},
			{Amount: decimal.NewFromFloat(-20.46)},
			{Amount: decimal.NewFromFloat(10)},
		}, decimal.NewFromFloat(39.74)},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, service.CalculateTotalBalance(tt.got))
		})
	}
}
