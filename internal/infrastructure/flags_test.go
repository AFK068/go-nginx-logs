package infrastructure_test

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestParseFlagToFlagConfigObject_success(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expected    *domain.FlagConfig
		expectError bool
	}{
		{
			name: "Valid flags",
			args: []string{
				"--path=some/path",
				"--from=2023-01-01T00:00:00Z",
				"--to=2023-12-31T23:59:59Z",
				"--format=markdown",
				"--filter-field=request",
				"--filter-value=GET",
			},
			expected: &domain.FlagConfig{
				Path:        "some/path",
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
				Format:      "markdown",
				FilterField: "request",
				FilterValue: "GET",
			},
			expectError: false,
		},
		{
			name: "Invalid format",
			args: []string{
				"--path=some/path",
				"--from=2023-01-01T00:00:00Z",
				"--to=2023-12-31T23:59:59Z",
				"--format=invalid",
				"--filter-field=request",
				"--filter-value=GET",
			},
			expected: &domain.FlagConfig{
				Path:        "some/path",
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
				Format:      "markdown",
				FilterField: "request",
				FilterValue: "GET",
			},
			expectError: false,
		},
		{
			name: "Empty from and to dates",
			args: []string{
				"--path=some/path",
				"--format=markdown",
				"--filter-field=request",
				"--filter-value=GET",
			},
			expected: &domain.FlagConfig{
				Path:        "some/path",
				From:        time.Time{},
				To:          time.Time{},
				Format:      "markdown",
				FilterField: "request",
				FilterValue: "GET",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)

			os.Args = append([]string{"cmd"}, tt.args...)

			result, err := infrastructure.ParseFlagToFlagConfigObject()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Path, result.Path)
				assert.True(t, tt.expected.From.Equal(result.From))
				assert.True(t, tt.expected.To.Equal(result.To))
				assert.Equal(t, tt.expected.Format, result.Format)
				assert.Equal(t, tt.expected.FilterField, result.FilterField)
				assert.Equal(t, tt.expected.FilterValue, result.FilterValue)
			}
		})
	}
}

func TestParseFlagToFlagConfigObject_failure(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expected    *domain.FlagConfig
		expectError bool
	}{
		{
			name: "Missing path flag",
			args: []string{
				"--from=2023-01-01T00:00:00Z",
				"--to=2023-12-31T23:59:59Z",
				"--format=markdown",
				"--filter-field=request",
				"--filter-value=GET",
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "Invalid date range",
			args: []string{
				"--path=some/path",
				"--from=2023-12-31T23:59:59Z",
				"--to=2023-01-01T00:00:00Z",
				"--format=markdown",
				"--filter-field=request",
				"--filter-value=GET",
			},
			expected:    nil,
			expectError: true,
		},
		{
			name: "To data before from date",
			args: []string{
				"--path=some/path",
				"--from=2023-01-01T00:00:00Z",
				"--to=2021-12-31T23:59:59Z",
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)

			os.Args = append([]string{"cmd"}, tt.args...)

			result, err := infrastructure.ParseFlagToFlagConfigObject()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Path, result.Path)
				assert.True(t, tt.expected.From.Equal(result.From))
				assert.True(t, tt.expected.To.Equal(result.To))
				assert.Equal(t, tt.expected.Format, result.Format)
				assert.Equal(t, tt.expected.FilterField, result.FilterField)
				assert.Equal(t, tt.expected.FilterValue, result.FilterValue)
			}
		})
	}
}
