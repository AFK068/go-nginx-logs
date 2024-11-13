package filter_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/filter"
	"github.com/stretchr/testify/assert"
)

func TestMatch_success(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		filterValue string
		expected    bool
	}{
		{
			name:        "Match string exact",
			value:       "test",
			filterValue: "test",
			expected:    true,
		},
		{
			name:        "Match string wildcard",
			value:       "test",
			filterValue: "te*",
			expected:    true,
		},
		{
			name:        "Match int exact",
			value:       123,
			filterValue: "123",
			expected:    true,
		},
		{
			name:        "Match time exact",
			value:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			filterValue: "2023-01-01T00:00:00Z",
			expected:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.Match(tt.value, tt.filterValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMatch_failure(t *testing.T) {
	tests := []struct {
		name        string
		value       any
		filterValue string
		expected    bool
	}{
		{
			name:        "No match string",
			value:       "test",
			filterValue: "nope",
			expected:    false,
		},
		{
			name:        "No match int",
			value:       123,
			filterValue: "321",
			expected:    false,
		},
		{
			name:        "No match time",
			value:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			filterValue: "2023-01-02T00:00:00Z",
			expected:    false,
		},
		{
			name:        "Invalid int filter value",
			value:       123,
			filterValue: "abc",
			expected:    false,
		},
		{
			name:        "Invalid time filter value",
			value:       time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			filterValue: "invalid-time",
			expected:    false,
		},
		{
			name:        "Unknown type",
			value:       map[int]int{},
			filterValue: "123",
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := filter.Match(tt.value, tt.filterValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}
