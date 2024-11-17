package integration_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestFlagConfig_FilterMatch_success(t *testing.T) {
	tests := []struct {
		name        string
		filterField string
		filterValue string
		log         *domain.NGINX
		expected    bool
	}{
		{
			name:        "Empty filter field and value",
			filterField: "",
			filterValue: "",
			log:         testLog,
			expected:    true,
		},
		{
			name:        "Empty filter value",
			filterField: "request",
			filterValue: "",
			log:         testLog,
			expected:    true,
		},
		{
			name:        "Filter match on request",
			filterField: "request",
			filterValue: "GET /downloads/product_1 HTTP/1.1",
			log:         testLog,
			expected:    true,
		},
		{
			name:        "Filter using *",
			filterField: "http_user_agent",
			filterValue: "Debian*",
			log:         testLog,
			expected:    true,
		},
		{
			name:        "Filter match on status",
			filterField: "status",
			filterValue: "304",
			log:         testLog,
			expected:    true,
		},
		{
			name:        "Filter match on time",
			filterField: "time_local",
			filterValue: "2023-05-17T08:05:32Z",
			log:         testLog,
			expected:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterConfig := &domain.FlagConfig{
				FilterField: tt.filterField,
				FilterValue: tt.filterValue,
			}

			result := filterConfig.FilterMatch(tt.log)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFlagConfig_FilterMatch_failure(t *testing.T) {
	tests := []struct {
		name        string
		filterField string
		filterValue string
		log         *domain.NGINX
		expected    bool
	}{
		{
			name:        "Incorrect filter field",
			filterField: "request",
			filterValue: "fail",
			log:         testLog,
			expected:    false,
		},
		{
			name:        "Filter no match on request",
			filterField: "request",
			filterValue: "POST /downloads/product_1 HTTP/1.1",
			log:         testLog,
			expected:    false,
		},
		{
			name:        "Filter no match on status",
			filterField: "status",
			filterValue: "200",
			log:         testLog,
			expected:    false,
		},
		{
			name:        "Filter no match on time",
			filterField: "time_local",
			filterValue: "2024-05-17T08:05:32Z",
			log:         testLog,
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterConfig := &domain.FlagConfig{
				FilterField: tt.filterField,
				FilterValue: tt.filterValue,
			}

			result := filterConfig.FilterMatch(tt.log)
			assert.Equal(t, tt.expected, result)
		})
	}
}
