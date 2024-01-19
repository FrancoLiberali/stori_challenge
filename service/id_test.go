package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTransactionID(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want uint
		err  error
	}{
		{"zero", "0", 0, nil},
		{"positive", "1", 1, nil},
		{"negative", "-1", 0, ErrParsingID},
		{"not convertible to int", "asd", 0, ErrParsingID},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			ans, err := parseTransactionID(tt.got)

			require.ErrorIs(t, err, tt.err)

			if err == nil {
				assert.Equal(t, tt.want, ans)
			}
		})
	}
}
