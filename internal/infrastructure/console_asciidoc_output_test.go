package infrastructure_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/streadway/quantile"
	"github.com/stretchr/testify/assert"
)

func TestAsciidocOutput(t *testing.T) {
	filterConfig := &domain.FlagConfig{
		From:        time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		To:          time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
		FilterField: "request",
		FilterValue: "GET /downloads/product_1 HTTP/1.1",
	}

	logReport := &domain.LogReport{
		FilterConfig:      filterConfig,
		NumberRequests:    1,
		TotalResponseSize: 100,
		MaxResponseSize:   100,
		MinResponseSize:   100,
		ResourceCount: map[string]int{
			"GET /downloads/product_1 HTTP/1.1": 1,
		},
		StatusCount: map[int]int{
			304: 1,
		},
		QuantileEstimator: quantile.New(quantile.Known(0.95, 0.001)),
	}

	logReport.QuantileEstimator.Add(100)

	var expectedOutput strings.Builder

	expectedOutput.WriteString("= General information\n\n")
	expectedOutput.WriteString("|===\n")
	expectedOutput.WriteString("| Metrics | Value\n")
	expectedOutput.WriteString(fmt.Sprintf("| Number of Requests | %d\n", logReport.NumberRequests))
	expectedOutput.WriteString(fmt.Sprintf("| Total Response Size | %d\n", logReport.TotalResponseSize))
	expectedOutput.WriteString(fmt.Sprintf("| Max Response Size | %d\n", logReport.MaxResponseSize))
	expectedOutput.WriteString(fmt.Sprintf("| Min Response Size | %d\n", logReport.MinResponseSize))
	expectedOutput.WriteString(fmt.Sprintf("| Average Response Size | %.2f\n", logReport.AverageResponseSize()))
	expectedOutput.WriteString(fmt.Sprintf("| 95th Percentile Response Size | %.2f\n", logReport.Percentile95()))
	expectedOutput.WriteString(fmt.Sprintf("| Start Date | %s\n", filterConfig.From.Format(time.RFC3339)))
	expectedOutput.WriteString(fmt.Sprintf("| End Date | %s\n", filterConfig.To.Format(time.RFC3339)))
	expectedOutput.WriteString(fmt.Sprintf("| Filter field | %s\n", filterConfig.FilterField))
	expectedOutput.WriteString(fmt.Sprintf("| Filter value | %s\n", filterConfig.FilterValue))
	expectedOutput.WriteString("|===\n\n")

	expectedOutput.WriteString("= Requested Resources\n\n")
	expectedOutput.WriteString("|===\n")
	expectedOutput.WriteString("| Resource | Count\n")
	expectedOutput.WriteString(fmt.Sprintf("| %s | %d\n", "GET /downloads/product_1 HTTP/1.1", 1))
	expectedOutput.WriteString("|===\n\n")

	expectedOutput.WriteString("= Response Codes\n\n")
	expectedOutput.WriteString("|===\n")
	expectedOutput.WriteString("| Code | Description | Count\n")
	expectedOutput.WriteString(fmt.Sprintf("| %d | %s | %d\n", 304, "Not Modified", 1))
	expectedOutput.WriteString("|===\n")
	expectedOutput.WriteString("\n")

	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	infrastructure.AsciidocOutput(logReport)
	w.Close()

	var actualOutput bytes.Buffer

	_, err := io.Copy(&actualOutput, r)
	assert.NoError(t, err)

	os.Stdout = stdout

	assert.Equal(t, expectedOutput.String(), actualOutput.String())
}
