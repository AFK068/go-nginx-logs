package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/nao1215/markdown"
)

func MarkdownOutput(logReport *domain.LogReport) {
	md := markdown.NewMarkdown(os.Stdout)

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

	md.H1("General information").Table(markdown.TableSet{
		Header: []string{"Metrics", "Value"},
		Rows: [][]string{
			{"Number of Requests", strconv.Itoa(logReport.NumberRequests)},
			{"Total Response Size", strconv.Itoa(logReport.TotalResponseSize)},
			{"Max Response Size", strconv.Itoa(logReport.MaxResponseSize)},
			{"Min Response Size", strconv.Itoa(logReport.MinResponseSize)},
			{"Average Response Size", fmt.Sprintf("%.2f", logReport.AverageResponseSize())},
			{"95th Percentile Response Size", fmt.Sprintf("%.2f", logReport.Percentile95())},
			{"Start Date", startDate},
			{"End Date", endDate},
			{"Filter field", filterField},
			{"Filter value", filterValue},
		},
	})

	md.H1("Requested Resources").Table(markdown.TableSet{
		Header: []string{"Resource", "Count"},
		Rows: func() [][]string {
			rows := [][]string{}
			for resource, count := range logReport.ResourceCount {
				rows = append(rows, []string{resource, strconv.Itoa(count)})
			}
			return rows
		}(),
	})

	md.H1("Response Codes").Table(markdown.TableSet{
		Header: []string{"Code", "Description", "Count"},
		Rows: func() [][]string {
			rows := [][]string{}
			for code, count := range logReport.StatusCount {
				rows = append(rows, []string{strconv.Itoa(code), http.StatusText(code), strconv.Itoa(count)})
			}
			return rows
		}(),
	})

	fmt.Println(md.String())
}
