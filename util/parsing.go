package util

import (
	"errors"
	"time"
)

func CalculateEndDate(startDate time.Time, durationValue int, durationUnit string) (time.Time, error) {
	switch durationUnit {
	case "day":
		return startDate.AddDate(0, 0, durationValue), nil
	case "month":
		return startDate.AddDate(0, durationValue, 0), nil
	case "year":
		return startDate.AddDate(durationValue, 0, 0), nil
	default:
		return time.Time{}, errors.New("invalid duration unit")
	}
}
