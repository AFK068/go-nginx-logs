package application

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/pathutils"
)

type Inizializer interface {
	InitializeConfig() (*domain.FlagConfig, error)
	InitializeLogReport(config *domain.FlagConfig) domain.LogReport
	InitializePath(paths string) (*pathutils.PathResult, error)
}

type DefaultInizializer struct{}

func (df *DefaultInizializer) InitializeConfig() (*domain.FlagConfig, error) {
	config, err := infrastructure.ParseFlagToFlagConfigObject()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (df *DefaultInizializer) InitializeLogReport(config *domain.FlagConfig) domain.LogReport {
	logReport := *domain.NewLogReport(config)

	return logReport
}

func (df *DefaultInizializer) InitializePath(paths string) (*pathutils.PathResult, error) {
	pathResult, err := pathutils.GetPath(paths)
	if err != nil {
		return nil, err
	}

	return pathResult, nil
}
