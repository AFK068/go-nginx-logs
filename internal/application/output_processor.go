package application

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

type OutputRenderer interface {
	Render(*domain.LogReport)
}

// Markdown render.
type MarkdownRenderer struct{}

func (mr *MarkdownRenderer) Render(logReport *domain.LogReport) {
	infrastructure.MarkdownOutput(logReport)
}

// Asciidoc render.
type AsciidocRenderer struct{}

func (ar *AsciidocRenderer) Render(logReport *domain.LogReport) {
	infrastructure.AsciidocOutput(logReport)
}

func GetOutputRenderer(outputFormat string) (OutputRenderer, error) {
	switch outputFormat {
	case domain.MarkdownFormat:
		return &MarkdownRenderer{}, nil
	case domain.AdocFormat:
		return &AsciidocRenderer{}, nil
	default:
		return nil, &InvalidRenderFormatError{Message: outputFormat}
	}
}
