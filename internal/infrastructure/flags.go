package infrastructure

import (
	"flag"
	"fmt"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

var (
	path   = flag.String("path", "", "Path to the file or URL (local path or URL)")
	from   = flag.String("from", "", "Start date in the format YYYY-MM-DDTHH:MM:SSZ (ISO8601)")
	to     = flag.String("to", "", "End date in the format YYYY-MM-DDTHH:MM:SSZ (ISO8601)")
	format = flag.String("format", "", "Output format, markdown or adoc")
	filter = flag.String("filter-value", "", "Filter parameter value")
)

func ParseFlagToFlagConfigObject() (*domain.FlagConfig, error) {
	flag.Parse()

	parsedPath := *path
	if isFlagEmpty(parsedPath) {
		return nil, &domain.MissedMandatoryFlagError{Message: "path"}
	}

	parsedFrom, err := parseDate(*from)
	if err != nil {
		fmt.Println("From flag has invalid format and will not be considered.")
	}

	parsedTo, err := parseDate(*to)
	if err != nil {
		fmt.Println("To flag has invalid format and will not be considered.")
	}

	parsedFormat := *format
	if isFlagEmpty(parsedFormat) {
		parsedFormat = domain.MarkdownFormat // Default format
	} else if parsedFormat != domain.MarkdownFormat && parsedFormat != domain.AdocFormat {
		fmt.Println("Format flag has invalid value and will be set to default value (markdown).")

		parsedFormat = domain.MarkdownFormat // Default format
	}

	parsedFilter := *filter
	if isFlagEmpty(parsedFilter) {
		parsedFilter = ""
	}

	return domain.NewFlagConfig(parsedPath, parsedFrom, parsedTo, parsedFormat, parsedFilter), nil
}

func isFlagEmpty(flagValue string) bool {
	return flagValue == ""
}

func parseDate(dateStr string) (time.Time, error) {
	if isFlagEmpty(dateStr) {
		return time.Time{}, nil
	}

	return time.Parse("2006-01-02", dateStr)
}
