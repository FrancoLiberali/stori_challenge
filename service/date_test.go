package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseDate(t *testing.T) {
	var tests = []struct {
		name string
		got  string
		want time.Time
		err  error
	}{
		{"month one digit, day one digit", "7/1", time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.UTC), nil},
		{"month two digits, day one digit", "11/2", time.Date(time.Now().Year(), 11, 2, 0, 0, 0, 0, time.UTC), nil},
		{"month two digits with zero, day one digit", "07/3", time.Date(time.Now().Year(), 7, 3, 0, 0, 0, 0, time.UTC), nil},
		{"month one digit, day two digits", "7/15", time.Date(time.Now().Year(), 7, 15, 0, 0, 0, 0, time.UTC), nil},
		{"month one digit, day two digits with zero", "7/04", time.Date(time.Now().Year(), 7, 4, 0, 0, 0, 0, time.UTC), nil},
		{"month two digits, day two digits", "11/12", time.Date(time.Now().Year(), 11, 12, 0, 0, 0, 0, time.UTC), nil},
		{"month two digits with zero, day two digits", "01/12", time.Date(time.Now().Year(), 1, 12, 0, 0, 0, 0, time.UTC), nil},
		{"month two digits, day two digits with zero", "10/02", time.Date(time.Now().Year(), 10, 2, 0, 0, 0, 0, time.UTC), nil},
		{"month two digits with zero, day two digits with zero", "01/02", time.Date(time.Now().Year(), 1, 2, 0, 0, 0, 0, time.UTC), nil},
		{"with year", "1/2/2023", time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), nil},
		{"only month", "2", time.Time{}, ErrParsingDate},
		{"month and day inverted", "15/2", time.Time{}, ErrParsingDate},
		{"parse not possible", "asd/asd", time.Time{}, ErrParsingDate},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			ans, err := parseDate(tt.got)

			require.ErrorIs(t, err, tt.err)

			if err == nil {
				assert.Equal(t, tt.want, ans)
			}
		})
	}
}
