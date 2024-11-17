package domain_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNGINX_GetFieldValue_success(t *testing.T) {
	nginx := domain.NewNGINX(
		"93.180.71.3",
		"-",
		time.Date(2015, 5, 17, 8, 5, 32, 0, time.FixedZone("UTC", 0)),
		"GET /downloads/product_1 HTTP/1.1",
		304,
		0,
		"-",
		"Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)",
	)

	tests := []struct {
		name     string
		field    string
		expected any
	}{
		{
			name:     "RemoteAddr",
			field:    "remote_addr",
			expected: "93.180.71.3",
		},
		{
			name:     "RemoteUser",
			field:    "remote_user",
			expected: "-",
		},
		{
			name:     "TimeLocal",
			field:    "time_local",
			expected: time.Date(2015, 5, 17, 8, 5, 32, 0, time.FixedZone("UTC", 0)),
		},
		{
			name:     "Request",
			field:    "request",
			expected: "GET /downloads/product_1 HTTP/1.1",
		},
		{
			name:     "Status",
			field:    "status",
			expected: 304,
		},
		{
			name:     "BodyBytesSent",
			field:    "body_bytes_sent",
			expected: 0,
		},
		{
			name:     "HTTPReferer",
			field:    "http_referer",
			expected: "-",
		},
		{
			name:     "HTTPUserAgent",
			field:    "http_user_agent",
			expected: "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := nginx.GetFieldValue(domain.NGINXFields(tt.field))
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNGINX_GetFieldValue_failure(t *testing.T) {
	nginx := domain.NewNGINX(
		"93.180.71.3",
		"-",
		time.Date(2015, 5, 17, 8, 5, 32, 0, time.FixedZone("UTC", 0)),
		"GET /downloads/product_1 HTTP/1.1",
		304,
		0,
		"-",
		"Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)",
	)

	tests := []struct {
		name     string
		field    string
		expected any
	}{
		{
			name:     "No find",
			field:    "error",
			expected: nil,
		},
		{
			name:     "Empty",
			field:    "",
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := nginx.GetFieldValue(domain.NGINXFields(tt.field))
			assert.Equal(t, tt.expected, result)
		})
	}
}
