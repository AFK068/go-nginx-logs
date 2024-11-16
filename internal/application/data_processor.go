package application

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/datastream"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/pathutils"
)

type DataProcessor interface {
	Process(*pathutils.PathResult, *domain.LogReport) error
}

// URL process.
type URLDataProcessor struct{}

func (p *URLDataProcessor) Process(paths *pathutils.PathResult, logReport *domain.LogReport) error {
	return datastream.ProcessFromURL(paths.Paths[0], &domain.NGINXParser{}, logReport)
}

// File process.
type FileDataProcessor struct{}

func (p *FileDataProcessor) Process(paths *pathutils.PathResult, logReport *domain.LogReport) error {
	return datastream.ProcessFromFile(paths.Paths, &domain.NGINXParser{}, logReport)
}

func GetDataProcessor(paths *pathutils.PathResult) (DataProcessor, error) {
	switch paths.Type {
	case "url":
		return &URLDataProcessor{}, nil
	case "file":
		return &FileDataProcessor{}, nil
	default:
		return nil, &InvalidPathFormatError{paths.Type}
	}
}
