package domain

import (
	"time"
)

var (
	MarkdownFormat = "markdown"
	AdocFormat     = "adoc"
)

type FlagConfig struct {
	Path   string
	From   time.Time
	To     time.Time
	Format string
	Filter string
}

func NewFlagConfig(path string, from, to time.Time, format, filter string) *FlagConfig {
	return &FlagConfig{path, from, to, format, filter}
}
