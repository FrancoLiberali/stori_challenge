package adapters

import (
	"encoding/csv"
	"fmt"
	"os"
)

type LocalCSVReader struct{}

// Read reads a CSV file by its name.
//
// Returns the list of rows of the CSV file
// or ErrReadingFile if an error is produced
func (reader LocalCSVReader) Read(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, err.Error())
	}

	// close the file after reading
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%w %s: %s", ErrReadingFile, fileName, err.Error())
	}

	// return error if file is empty
	if len(data) <= 1 {
		return nil, fmt.Errorf("%w %s: file has less that 2 lines", ErrReadingFile, fileName)
	}

	// avoid returning the header
	return data[1:], nil
}
