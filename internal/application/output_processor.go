package application

import (
	"log/slog"

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

	slog.Info("markdown output rendered")
}

// Asciidoc render.
type AsciidocRenderer struct{}

func (ar *AsciidocRenderer) Render(logReport *domain.LogReport) {
	infrastructure.AsciidocOutput(logReport)

	slog.Info("asciidoc output rendered")
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
