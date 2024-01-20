package service

import (
	"errors"
	"fmt"
	"strconv"
)

var ErrParsingID = errors.New("can't convert transaction id")

// parseTransactionID transforms a string representing an id to uint
//
// Returns ErrParsingID if the transformation if not possible
func parseTransactionID(idString string) (uint, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrParsingID, err.Error())
	}

	if id < 0 {
		return 0, fmt.Errorf("%w: %s", ErrParsingID, "value cannot be less than 0")
	}

	return uint(id), nil
}
