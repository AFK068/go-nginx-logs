package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestGetDataProcessor(t *testing.T) {
	tests := []struct {
		Name             string
		PathResultConfig pathutils.PathResult
		expectedType     interface{}
		expectError      bool
	}{
		{
			Name:             "Valid url path type",
			PathResultConfig: pathutils.PathResult{Type: "url"},
			expectedType:     &application.URLDataProcessor{},
			expectError:      false,
		},
		{
			Name:             "Valid path type",
			PathResultConfig: pathutils.PathResult{Type: "file"},
			expectedType:     &application.FileDataProcessor{},
			expectError:      false,
		},
		{
			Name:             "Invalid format",
			PathResultConfig: pathutils.PathResult{Type: "invalid"},
			expectedType:     nil,
			expectError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			processor, err := application.GetDataProcessor(&tt.PathResultConfig)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, processor)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, tt.expectedType, processor)
			}
		})
	}
}
