package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseAmount(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want float64
		err  error
	}{
		{"positive one decimal", "+60.5", 60.5, nil},
		{"negative one decimal", "-10.3", -10.3, nil},
		{"positive two decimals", "23.48", 23.48, nil},
		{"negative two decimals", "-20.46", -20.46, nil},
		{"positive no decimals", "+10", 10, nil},
		{"negative no decimals", "-10", -10, nil},
		{"not convertible to float", "asd", 0, ErrParsingAmount},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			ans, err := parseAmount(tt.got)

			require.ErrorIs(t, err, tt.err)

			if err == nil {
				float, _ := ans.Float64()
				assert.InEpsilon(t, float, tt.want, 0)
			}
		})
	}
}
