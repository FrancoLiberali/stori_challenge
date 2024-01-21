package adapters

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocalCSVReader(t *testing.T) {
	reader := LocalCSVReader{}

	tests := []struct {
		setup      func(t *testing.T)
		name       string
		got        string
		want       [][]string
		err        error
		errMessage string
		teardown   func(t *testing.T)
	}{
		{
			func(t *testing.T) {},
			"file not found",
			"not_found.csv",
			nil,
			ErrReadingFile,
			"not_found.csv: open not_found.csv: no such file or directory",
			func(t *testing.T) {},
		},
		{
			func(t *testing.T) {
				writeFile(t, "found.txt", `asd,
asd`)
			},
			"file not csv",
			"found.txt",
			nil,
			ErrReadingFile,
			"found.txt: record on line 2: wrong number of fields",
			func(t *testing.T) { removeFile(t, "found.txt") },
		},
		{
			func(t *testing.T) { writeFile(t, "found.csv", "Id,Date,Transaction") },
			"file not enough lines",
			"found.csv",
			nil,
			ErrReadingFile,
			"found.csv: file has less that 2 lines",
			func(t *testing.T) { removeFile(t, "found.csv") },
		},
		{
			func(t *testing.T) {
				writeFile(t, "correct.csv", `Id,Date,Transaction
0,7/15,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10`)
			},
			"correct file",
			"correct.csv",
			[][]string{
				{"0", "7/15", "+60.5"},
				{"1", "7/28", "-10.3"},
				{"2", "8/2", "-20.46"},
				{"3", "8/13", "+10"},
			},
			nil,
			"",
			func(t *testing.T) { removeFile(t, "correct.csv") },
		},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)

			ans, err := reader.Read(tt.got)

			tt.teardown(t)

			require.ErrorIs(t, err, tt.err)

			if err == nil {
				assert.ElementsMatch(t, ans, tt.want)
			} else {
				require.ErrorContains(t, err, tt.errMessage)
			}
		})
	}
}

func writeFile(t *testing.T, fileName string, content string) {
	file, err := os.Create(fileName)
	require.NoError(t, err)
	_, err = file.WriteString(content)
	require.NoError(t, err)
	err = file.Close()
	require.NoError(t, err)
}

func removeFile(t *testing.T, fileName string) {
	err := os.Remove(fileName)
	require.NoError(t, err)
}
