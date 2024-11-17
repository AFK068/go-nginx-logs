package domain_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNGINXParse_success(t *testing.T) {
	tests := []struct {
		name        string
		logString   string
		expected    *domain.NGINX
		expectError bool
	}{
		{
			name:      "Valid log string",
			logString: `93.180.71.3 - - [17/May/2015:08:05:32 +0000] "GET /downloads/product_1 HTTP/1.1" 304 0 "-" "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)"`, //nolint:lll // large strign
			expected: &domain.NGINX{
				RemoteAddr:    "93.180.71.3",
				RemoteUser:    "-",
				TimeLocal:     time.Date(2015, 5, 17, 8, 5, 32, 0, time.FixedZone("UTC", 0)),
				Request:       "GET /downloads/product_1 HTTP/1.1",
				Status:        304,
				BodyBytesSent: 0,
				HTTPReferer:   "-",
				HTTPUserAgent: "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &domain.NGINXParser{}
			result, err := parser.Parse(tt.logString)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.RemoteAddr, result.RemoteAddr)
				assert.Equal(t, tt.expected.RemoteUser, result.RemoteUser)
				assert.True(t, tt.expected.TimeLocal.Equal(result.TimeLocal))
				assert.Equal(t, tt.expected.Request, result.Request)
				assert.Equal(t, tt.expected.Status, result.Status)
				assert.Equal(t, tt.expected.BodyBytesSent, result.BodyBytesSent)
				assert.Equal(t, tt.expected.HTTPReferer, result.HTTPReferer)
				assert.Equal(t, tt.expected.HTTPUserAgent, result.HTTPUserAgent)
			}
		})
	}
}

func TestNGINXParse_failure(t *testing.T) {
	tests := []struct {
		name        string
		logString   string
		expected    *domain.NGINX
		expectError bool
	}{

		{
			name:        "Invalid format '-' not found",
			logString:   `93.180.71.3 - [17/May/2015:08:05:32 +0000] "GET /downloads/product_1 HTTP/1.1" 304 0 "-" "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)"`, //nolint:lll // large strign
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Invalid timeLocal format",
			logString:   `93.180.71.3 - - [27-10-2005] "GET /downloads/product_1 HTTP/1.1" 304 0 "-" "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)"`, //nolint:lll // large strign
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Invalid status format",
			logString:   `93.180.71.3 - - [17/May/2015:08:05:32 +0000] "GET /downloads/product_1 HTTP/1.1" invalid 0 "-" "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)"`, //nolint:lll // large strign
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Invalid bodyBytesSent format",
			logString:   `93.180.71.3 - - [17/May/2015:08:05:32 +0000] "GET /downloads/product_1 HTTP/1.1" 304 invalid "-" "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)"`, //nolint:lll // large strign
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &domain.NGINXParser{}
			result, err := parser.Parse(tt.logString)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
