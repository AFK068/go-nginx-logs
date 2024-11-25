package integration_test

import (
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/streadway/quantile"
	"github.com/stretchr/testify/assert"
)

var testLog = &domain.NGINX{
	RemoteAddr:    "93.180.71.3",
	RemoteUser:    "-",
	TimeLocal:     time.Date(2023, 5, 17, 8, 5, 32, 0, time.UTC),
	Request:       "GET /downloads/product_1 HTTP/1.1",
	Status:        304,
	BodyBytesSent: 0,
	HTTPReferer:   "-",
	HTTPUserAgent: "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)",
}

func TestLogReport_Update(t *testing.T) {
	tests := []struct {
		name         string
		filterConfig *domain.FlagConfig
		log          *domain.NGINX
		expected     *domain.LogReport
	}{
		{
			name: "Valid log string",
			filterConfig: &domain.FlagConfig{
				From: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				NumberRequests:    1,
				TotalResponseSize: 0,
				MaxResponseSize:   0,
				MinResponseSize:   0,
				ResourceCount: map[string]int{
					"GET /downloads/product_1 HTTP/1.1": 1,
				},
				StatusCount: map[int]int{
					304: 1,
				},
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
		{
			name: "Log out of range",
			filterConfig: &domain.FlagConfig{
				From: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2023, 2, 1, 0, 0, 0, 0, time.UTC),
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				MaxResponseSize:   0,
				MinResponseSize:   int(^uint(0) >> 1),
				ResourceCount:     make(map[string]int),
				StatusCount:       make(map[int]int),
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
		{
			name: "Filter match",
			filterConfig: &domain.FlagConfig{
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
				FilterField: "request",
				FilterValue: "GET /downloads/product_1 HTTP/1.1",
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				NumberRequests:    1,
				TotalResponseSize: 0,
				MaxResponseSize:   0,
				MinResponseSize:   0,
				ResourceCount: map[string]int{
					"GET /downloads/product_1 HTTP/1.1": 1,
				},
				StatusCount: map[int]int{
					304: 1,
				},
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
		{
			name: "Filter no match",
			filterConfig: &domain.FlagConfig{
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
				FilterField: "fail",
				FilterValue: "test",
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				MaxResponseSize:   0,
				MinResponseSize:   int(^uint(0) >> 1),
				ResourceCount:     make(map[string]int),
				StatusCount:       make(map[int]int),
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
		{
			name: "Filters empty",
			filterConfig: &domain.FlagConfig{
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
				FilterField: "",
				FilterValue: "",
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				NumberRequests:    1,
				TotalResponseSize: 0,
				MaxResponseSize:   0,
				MinResponseSize:   0,
				ResourceCount: map[string]int{
					"GET /downloads/product_1 HTTP/1.1": 1,
				},
				StatusCount: map[int]int{
					304: 1,
				},
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
		{
			name: "Filter with *",
			filterConfig: &domain.FlagConfig{
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
				FilterField: "request",
				FilterValue: "GET*",
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				NumberRequests:    1,
				TotalResponseSize: 0,
				MaxResponseSize:   0,
				MinResponseSize:   0,
				ResourceCount: map[string]int{
					"GET /downloads/product_1 HTTP/1.1": 1,
				},
				StatusCount: map[int]int{
					304: 1,
				},
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
		{
			name: "Incorrect time format filter",
			filterConfig: &domain.FlagConfig{
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
				FilterField: "time_local",
				FilterValue: "20023-05-10",
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				MaxResponseSize:   0,
				MinResponseSize:   int(^uint(0) >> 1),
				ResourceCount:     make(map[string]int),
				StatusCount:       make(map[int]int),
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
		{
			name: "Correct time format filter",
			filterConfig: &domain.FlagConfig{
				From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
				To:          time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC),
				FilterField: "time_local",
				FilterValue: "2023-05-17T08:05:32Z",
			},
			log: testLog,
			expected: &domain.LogReport{
				FilterConfig:      &domain.FlagConfig{},
				NumberRequests:    1,
				TotalResponseSize: 0,
				MaxResponseSize:   0,
				MinResponseSize:   0,
				ResourceCount: map[string]int{
					"GET /downloads/product_1 HTTP/1.1": 1,
				},
				StatusCount: map[int]int{
					304: 1,
				},
				QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logReport := domain.NewLogReport(tt.filterConfig)

			logReport.Update(tt.log)

			assert.Equal(t, tt.expected.NumberRequests, logReport.NumberRequests)
			assert.Equal(t, tt.expected.TotalResponseSize, logReport.TotalResponseSize)
			assert.Equal(t, tt.expected.MaxResponseSize, logReport.MaxResponseSize)
			assert.Equal(t, tt.expected.MinResponseSize, logReport.MinResponseSize)
			assert.Equal(t, tt.expected.ResourceCount, logReport.ResourceCount)
			assert.Equal(t, tt.expected.StatusCount, logReport.StatusCount)
			assert.Equal(t, tt.expected.QuantileEstimator.Get(0.95), logReport.QuantileEstimator.Get(0.95))
		})
	}
}
