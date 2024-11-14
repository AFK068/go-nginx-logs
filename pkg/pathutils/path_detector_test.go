package pathutils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestGetPath_success(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdata")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	tempFile1, err := os.CreateTemp(tempDir, "file1.txt")
	assert.NoError(t, err)

	tempFile2, err := os.CreateTemp(tempDir, "file2.txt")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		input       string
		expected    *pathutils.PathResult
		expectError bool
	}{
		{
			name:  "Valid URL",
			input: "https://example.com",
			expected: &pathutils.PathResult{
				Paths: []string{"https://example.com"},
				Type:  "url",
			},
			expectError: false,
		},
		{
			name:  "Valid local path",
			input: filepath.Join(tempDir, "file*"),
			expected: &pathutils.PathResult{
				Paths: []string{tempFile1.Name(), tempFile2.Name()},
				Type:  "file",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := pathutils.GetPath(tt.input)
			if tt.expectError {
				assert.Error(t, err)
				assert.IsType(t, &pathutils.PathError{}, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetPath_failure(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "testdata")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name        string
		input       string
		expected    *pathutils.PathResult
		expectError bool
	}{
		{
			name:  "Invalid URL",
			input: "htp://invalid-url",
			expected: &pathutils.PathResult{
				Paths: nil,
				Type:  "",
			},
			expectError: true,
		},
		{
			name:  "Non-existent local path",
			input: filepath.Join(tempDir, "nonexistent"),
			expected: &pathutils.PathResult{
				Paths: nil,
				Type:  "",
			},
			expectError: true,
		},
		{
			name:  "Directory path",
			input: tempDir,
			expected: &pathutils.PathResult{
				Paths: nil,
				Type:  "",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := pathutils.GetPath(tt.input)
			if tt.expectError {
				assert.Error(t, err)
				assert.IsType(t, &pathutils.PathError{}, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
