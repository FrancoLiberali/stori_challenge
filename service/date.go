package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	receivedDateSeparator = "/"
	wantDateSeparator     = "-"
	dateLayout            = "07/15/2024"
)

var ErrParsingDate = errors.New("can't convert transaction date")

func parseDate(dateString string) (time.Time, error) {
	dateSplitted := strings.Split(dateString, receivedDateSeparator)
	if len(dateSplitted) < 2 {
		return time.Time{}, fmt.Errorf("%w: at least month and year are expected", ErrParsingDate)
	}

	// add a 0 to the begging of the date to ensure it has to digits
	for i, datePart := range dateSplitted {
		if len(datePart) == 1 {
			dateSplitted[i] = "0" + datePart
		}
	}

	// if dateString doesn't have the year, assume this year
	if len(dateSplitted) == 2 {
		dateSplitted = append(dateSplitted, strconv.Itoa(time.Now().Year()))
	}

	dateString = dateSplitted[2] + wantDateSeparator + dateSplitted[0] + wantDateSeparator + dateSplitted[1]

	ans, err := time.Parse(time.DateOnly, dateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %s", ErrParsingDate, err.Error())
	}

	return ans, nil
}
