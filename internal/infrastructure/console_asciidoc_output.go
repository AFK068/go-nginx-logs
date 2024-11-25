package infrastructure

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func AsciidocOutput(logReport *domain.LogReport) {
	var sb strings.Builder

	startDate := "-"
	if !logReport.FilterConfig.From.IsZero() {
		startDate = logReport.FilterConfig.From.Format(time.RFC3339)
	}

	endDate := "-"
	if !logReport.FilterConfig.To.IsZero() {
		endDate = logReport.FilterConfig.To.Format(time.RFC3339)
	}

	filterField := "-"
	if logReport.FilterConfig.FilterField != "" {
		filterField = logReport.FilterConfig.FilterField
	}

	filterValue := "-"
	if logReport.FilterConfig.FilterValue != "" {
		filterValue = logReport.FilterConfig.FilterValue
	}

	if logReport.MinResponseSize == int(^uint(0)>>1) && logReport.NumberRequests == 0 {
		logReport.MinResponseSize = 0
	}

	sb.WriteString("= General information\n\n")
	sb.WriteString("|===\n")
	sb.WriteString("| Metrics | Value\n")
	sb.WriteString(fmt.Sprintf("| Number of Requests | %d\n", logReport.NumberRequests))
	sb.WriteString(fmt.Sprintf("| Total Response Size | %d\n", logReport.TotalResponseSize))
	sb.WriteString(fmt.Sprintf("| Max Response Size | %d\n", logReport.MaxResponseSize))
	sb.WriteString(fmt.Sprintf("| Min Response Size | %d\n", logReport.MinResponseSize))
	sb.WriteString(fmt.Sprintf("| Average Response Size | %.2f\n", logReport.AverageResponseSize()))
	sb.WriteString(fmt.Sprintf("| 95th Percentile Response Size | %.2f\n", logReport.Percentile95()))
	sb.WriteString(fmt.Sprintf("| Start Date | %s\n", startDate))
	sb.WriteString(fmt.Sprintf("| End Date | %s\n", endDate))
	sb.WriteString(fmt.Sprintf("| Filter field | %s\n", filterField))
	sb.WriteString(fmt.Sprintf("| Filter value | %s\n", filterValue))
	sb.WriteString("|===\n\n")

	sb.WriteString("= Requested Resources\n\n")
	sb.WriteString("|===\n")
	sb.WriteString("| Resource | Count\n")

	for resource, count := range logReport.ResourceCount {
		sb.WriteString(fmt.Sprintf("| %s | %d\n", resource, count))
	}

	sb.WriteString("|===\n\n")

	sb.WriteString("= Response Codes\n\n")
	sb.WriteString("|===\n")
	sb.WriteString("| Code | Description | Count\n")

	for code, count := range logReport.StatusCount {
		sb.WriteString(fmt.Sprintf("| %d | %s | %d\n", code, http.StatusText(code), count))
	}

	sb.WriteString("|===\n")

	fmt.Println(sb.String())
}
