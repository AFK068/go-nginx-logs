package pathutils

import (
	"os"
	"path/filepath"

	urlverifier "github.com/davidmytton/url-verifier"
)

// Determines the URL of a string or a local path and gets the corresponding results.
func GetPath(path string) ([]string, error) {
	// Check for local path.
	files, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}

	// Filtering directories.
	var validFiles []string

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return nil, err
		}

		if !info.IsDir() {
			validFiles = append(validFiles, file)
		}
	}

	if len(validFiles) != 0 {
		return validFiles, nil
	}

	_, err = urlverifier.NewVerifier().CheckHTTP(path)
	if err != nil {
		return nil, err
	}

	return []string{path}, nil
}
