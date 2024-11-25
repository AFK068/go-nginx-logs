package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetOutputRenderer(t *testing.T) {
	tests := []struct {
		name         string
		outputFormat string
		expectedType interface{}
		expectError  bool
	}{
		{
			name:         "Valid markdown format",
			outputFormat: domain.MarkdownFormat,
			expectedType: &application.MarkdownRenderer{},
			expectError:  false,
		},
		{
			name:         "Valid adoc format",
			outputFormat: domain.AdocFormat,
			expectedType: &application.AsciidocRenderer{},
			expectError:  false,
		},
		{
			name:         "Invalid format",
			outputFormat: "invalid",
			expectedType: nil,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer, err := application.GetOutputRenderer(tt.outputFormat)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, renderer)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, tt.expectedType, renderer)
			}
		})
	}
}
