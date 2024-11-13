package timeutils_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/timeutils"
	"github.com/stretchr/testify/assert"
)

func TestInTimeSpan(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		check    time.Time
		expected bool
	}{
		{
			name:     "Within range",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Outside range",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 2, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Start equals check",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "End equals check",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "End before start",
			start:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Custom format within range",
			start:    parseCustomTime("17/May/2015:00:00:00 +0000"),
			end:      parseCustomTime("18/May/2015:00:00:00 +0000"),
			check:    parseCustomTime("17/May/2015:08:05:32 +0000"),
			expected: true,
		},
		{
			name:     "Custom format outside range",
			start:    parseCustomTime("17/May/2015:00:00:00 +0000"),
			end:      parseCustomTime("18/May/2015:00:00:00 +0000"),
			check:    parseCustomTime("18/May/2015:08:05:32 +0000"),
			expected: false,
		},
		{
			name:     "Empty start and end",
			start:    time.Time{},
			end:      time.Time{},
			check:    time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Check is zero time",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			check:    time.Time{},
			expected: false,
		},
		{
			name:     "Check equals start",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Check equals end",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "Check equals start and end",
			start:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			check:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := timeutils.InTimeSpan(tt.start, tt.end, tt.check)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func parseCustomTime(value string) time.Time {
	t, _ := time.Parse("02/Jan/2006:15:04:05 -0700", value)
	return t
}
