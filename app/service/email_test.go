package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionsPerMonthToEmailData(t *testing.T) {
	tests := []struct {
		name string
		got  []TransactionsPerMonth
		want []transactionsPerMonthEmailData
	}{
		{"empty list", []TransactionsPerMonth{}, []transactionsPerMonthEmailData{}},
		{
			"list with 1 transaction in one month",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 1},
			},
			[]transactionsPerMonthEmailData{{Month: "July", Value: "1", IsOdd: true}},
		},
		{
			"list with multiple transaction in one month",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{{Month: "July", Value: "2", IsOdd: true}},
		},
		{
			"list with multiple months",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(time.Now().Year(), 8, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{
				{Month: "July", Value: "2", IsOdd: true},
				{Month: "August", Value: "2", IsOdd: false},
			},
		},
		{
			"list with multiple months of another year",
			[]TransactionsPerMonth{
				{Month: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{
				{Month: "July 2023", Value: "2", IsOdd: true},
				{Month: "August 2023", Value: "2", IsOdd: false},
			},
		},
		{
			"list with multiple years",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
			},
			[]transactionsPerMonthEmailData{
				{Month: "July", Value: "2", IsOdd: true},
				{Month: "July 2023", Value: "2", IsOdd: false},
			},
		},
		{
			"list 3 months",
			[]TransactionsPerMonth{
				{Month: time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), Amount: 2},
				{Month: time.Date(time.Now().Year(), 8, 1, 0, 0, 0, 0, time.UTC), Amount: 3},
				{Month: time.Date(time.Now().Year(), 9, 1, 0, 0, 0, 0, time.UTC), Amount: 4},
			},
			[]transactionsPerMonthEmailData{
				{Month: "July", Value: "2", IsOdd: true},
				{Month: "August", Value: "3", IsOdd: false},
				{Month: "September", Value: "4", IsOdd: true},
			},
		},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, transactionsPerMonthToEmailData(tt.got))
		})
	}
}
