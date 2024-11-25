package pathutils

import (
	"os"
	"path/filepath"

	urlverifier "github.com/davidmytton/url-verifier"
)

// PathResult contains the results and type of the path.
type PathResult struct {
	Paths []string
	Type  string
}

// GetPath determines whether the string is a URL or a local path and returns the corresponding results and path type.
func GetPath(path string) (*PathResult, error) {
	pathFiles, err := getFile(path)
	if err != nil {
		return nil, &PathError{Message: err.Error()}
	}

	if len(pathFiles) > 0 {
		return &PathResult{Paths: pathFiles, Type: "file"}, nil
	}

	pathURL, err := getURL(path)
	if err != nil {
		return nil, &PathError{Message: err.Error()}
	}

	return &PathResult{Paths: pathURL, Type: "url"}, nil
}

// getFile checks if the string is a local path and returns the corresponding results.
func getFile(path string) ([]string, error) {
	files, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}

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

	return validFiles, nil
}

// getURL checks if the string is a URL and returns the corresponding results.
func getURL(path string) ([]string, error) {
	_, err := urlverifier.NewVerifier().CheckHTTP(path)
	if err != nil {
		return nil, err
	}

	return []string{path}, nil
}
