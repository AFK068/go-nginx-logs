package filter

import (
	"strconv"
	"strings"
	"time"
)

func Match(value any, filterValue string) bool {
	switch v := value.(type) {
	case string:
		return matchString(v, filterValue)
	case int:
		return matchInt(v, filterValue)
	case time.Time:
		return matchTime(v, filterValue)
	default:
		return false
	}
}

func matchString(value, filterValue string) bool {
	if strings.Contains(filterValue, "*") {
		return strings.Contains(value, strings.ReplaceAll(filterValue, "*", ""))
	}

	return value == filterValue
}

func matchInt(value int, filterValue string) bool {
	filterInt, err := strconv.Atoi(filterValue)
	if err != nil {
		return false
	}

	return value == filterInt
}

func matchTime(value time.Time, filterValue string) bool {
	timeParsed, err := time.Parse(time.RFC3339, filterValue)
	if err != nil {
		return false
	}

	return value.Equal(timeParsed)
}
