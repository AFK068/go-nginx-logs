package timeutils

import "time"

func InTimeSpan(start, end, check time.Time) bool {
	if start.IsZero() && end.IsZero() {
		return true
	}

	if start.Equal(end) {
		return check.Equal(start)
	}

	if end.Before(start) {
		return (check.After(start) || check.Equal(start)) || (check.Before(end) || check.Equal(end))
	}

	return (check.After(start) || check.Equal(start)) && (check.Before(end) || check.Equal(end))
}
