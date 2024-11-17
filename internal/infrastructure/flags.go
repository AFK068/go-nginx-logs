package infrastructure

import (
	"flag"
	"fmt"
	"log/slog"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
)

func ParseFlagToFlagConfigObject() (*domain.FlagConfig, error) {
	var (
		path        = flag.String("path", "", "Path to the file or URL (local path or URL)")
		from        = flag.String("from", "", "Start date in the format YYYY-MM-DDTHH:MM:SSZ (ISO8601)")
		to          = flag.String("to", "", "End date in the format YYYY-MM-DDTHH:MM:SSZ (ISO8601)")
		format      = flag.String("format", "", "Output format, markdown or adoc")
		filterField = flag.String("filter-field", "", "Filter parameter field")
		filterValue = flag.String("filter-value", "", "Filter parameter value")
	)

	flag.Parse()

	parsedPath := *path
	if isFlagEmpty(parsedPath) {
		slog.Error("missing mandatory flag")

		return nil, &domain.MissedMandatoryFlagError{Message: "path"}
	}

	parsedFrom, err := parseDate(*from)
	if err != nil {
		slog.Error("failed to parse date", slog.String("error", err.Error()))

		fmt.Println("From flag has invalid format and will not be considered.")
	}

	parsedTo, err := parseDate(*to)
	if err != nil {
		slog.Error("failed to parse date", slog.String("error", err.Error()))

		fmt.Println("To flag has invalid format and will not be considered.")
	}

	if parsedTo.Before(parsedFrom) && !parsedTo.IsZero() && !parsedFrom.IsZero() {
		slog.Error("end date is before start date")

		return nil, &domain.InvalidDateRangeError{Message: "end date is before start date"}
	}

	parsedFormat := *format
	if isFlagEmpty(parsedFormat) {
		parsedFormat = domain.MarkdownFormat // Default format

		slog.Info("format flag is empty and will be set to default value (markdown)")
	} else if parsedFormat != domain.MarkdownFormat && parsedFormat != domain.AdocFormat {
		fmt.Println("Format flag has invalid value and will be set to default value (markdown).")
		slog.Info("format flag has invalid value and set to default value (markdown)")

		parsedFormat = domain.MarkdownFormat // Default format
	}

	parsedFilterField := *filterField
	if isFlagEmpty(parsedFilterField) {
		slog.Info("filter field is empty and will not be considered")

		parsedFilterField = ""
	}

	parsedFilterValue := *filterValue
	if isFlagEmpty(parsedFilterValue) {
		slog.Info("filter value is empty and will not be considered")

		parsedFilterValue = ""
	}

	slog.Info("flag config object initialized successfully: ",
		slog.String("path", parsedPath), slog.String("from", parsedFrom.String()),
		slog.String("to", parsedTo.String()), slog.String("format", parsedFormat),
		slog.String("filterField", parsedFilterField), slog.String("filterValue", parsedFilterValue))

	return domain.NewFlagConfig(parsedPath, parsedFrom, parsedTo, parsedFormat, parsedFilterField, parsedFilterValue), nil
}

func isFlagEmpty(flagValue string) bool {
	return flagValue == ""
}

func parseDate(dateStr string) (time.Time, error) {
	if isFlagEmpty(dateStr) {
		return time.Time{}, nil
	}

	return time.Parse(time.RFC3339, dateStr)
}
