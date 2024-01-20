package adapters

import (
	"errors"
)

type CSVReader interface {
	// Read reads a CSV file by its name.
	//
	// Returns the list of rows of the CSV file
	// or ErrReadingFile if an error is produced
	Read(fileName string) ([][]string, error)
}

var ErrReadingFile = errors.New("error while reading file")
