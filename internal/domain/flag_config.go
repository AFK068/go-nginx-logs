package domain

import "time"

var (
	MarkdownFormat = "markdown"
	AdocFormat     = "adoc"
)

type FlagConfig struct {
	path   string
	from   time.Time
	to     time.Time
	format string
	filter string
}

func NewFlagConfig(path string, from, to time.Time, format, filter string) *FlagConfig {
	return &FlagConfig{path, from, to, format, filter}
}
