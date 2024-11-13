package domain

import (
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/filter"
)

var (
	MarkdownFormat = "markdown"
	AdocFormat     = "adoc"
	Empty          = ""
)

type FlagConfig struct {
	Path        string
	From        time.Time
	To          time.Time
	Format      string
	FilterField string
	FilterValue string
}

func NewFlagConfig(path string, from, to time.Time, format, filterField, filterValue string) *FlagConfig {
	return &FlagConfig{path, from, to, format, filterField, filterValue}
}

func (fc *FlagConfig) FilterMatch(log *NGINX) bool {
	if fc.FilterField == Empty {
		return true
	}

	value := log.GetFieldValue(NGINXFields(fc.FilterField))
	if value == nil {
		return false
	}

	if fc.FilterValue == Empty {
		return true
	}

	return filter.Match(value, fc.FilterValue)
}
