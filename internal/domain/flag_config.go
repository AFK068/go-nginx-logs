package domain

import (
	"time"
)

var (
	MarkdownFormat = "markdown"
	AdocFormat     = "adoc"
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
