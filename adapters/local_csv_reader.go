package adapters

import (
	"encoding/csv"
	"fmt"
	"os"
)

type LocalCsvReader struct{}

// Read reads a CSV file by its name.
//
// Returns the list of rows of the CSV file
// or ErrReadingFile if an error is produced
func (reader LocalCsvReader) Read(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, err.Error())
	}

	// remember to close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, err.Error())
	}

	if len(data) <= 1 {
		return nil, fmt.Errorf("%w %s: file has less that 2 lines", ErrReadingFile, fileName)
	}

	return data[1:], nil
}
