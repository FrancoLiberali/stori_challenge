package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionsPerMonthToString(t *testing.T) {
	tests := []struct {
		name string
		got  []TransactionsPerMonth
		want string
	}{
		{"empty list", []TransactionsPerMonth{}, ""},
		{"list with 1 transaction in one month", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 1},
		}, `Number of transactions in July: 1
`},
		{"list with multiple transaction in one month", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July: 2
`},
		{"list with multiple months", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			{Month: time.Date(time.Now().Year(), 8, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July: 2
Number of transactions in August: 2
`},
		{"list with multiple months of another year", []TransactionsPerMonth{
			{Month: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			{Month: time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July 2023: 2
Number of transactions in August 2023: 2
`},
		{"list with multiple years", []TransactionsPerMonth{
			{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			{Month: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
		}, `Number of transactions in July: 2
Number of transactions in July 2023: 2
`},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, transactionsPerMonthToString(tt.got))
		})
	}
}
