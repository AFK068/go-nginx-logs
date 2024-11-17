package domain_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestLogReport_AverageResponseSize(t *testing.T) {
	tests := []struct {
		name                string
		numberRequests      int
		totalResponseSize   int
		expectedAverageSize float64
	}{
		{
			name:                "Average response size with 100 requests",
			numberRequests:      100,
			totalResponseSize:   1000,
			expectedAverageSize: 10,
		},
		{
			name:                "Average response size with 50 requests",
			numberRequests:      50,
			totalResponseSize:   500,
			expectedAverageSize: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logReport := domain.NewLogReport(nil)
			logReport.NumberRequests = tt.numberRequests
			logReport.TotalResponseSize = tt.totalResponseSize

			assert.Equal(t, tt.expectedAverageSize, logReport.AverageResponseSize())
		})
	}
}

func TestLogReport_Percentile95(t *testing.T) {
	tests := []struct {
		name               string
		numberRequests     int
		values             []float64
		expectedPercentile float64
	}{
		{
			name:               "Percentile 95 with 10 values",
			numberRequests:     10,
			values:             []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedPercentile: 9,
		},
		{
			name:               "Percentile 95 with 5 values",
			numberRequests:     5,
			values:             []float64{10, 20, 30, 40, 50},
			expectedPercentile: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logReport := domain.NewLogReport(nil)
			logReport.NumberRequests = tt.numberRequests

			for _, value := range tt.values {
				logReport.QuantileEstimator.Add(value)
			}

			assert.Equal(t, float64(tt.expectedPercentile), logReport.Percentile95())
		})
	}
}
