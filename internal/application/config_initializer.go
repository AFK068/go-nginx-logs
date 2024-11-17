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
	return infrastructure.ParseFlagToFlagConfigObject()
}

func (df *DefaultInizializer) InitializeLogReport(config *domain.FlagConfig) domain.LogReport {
	return *domain.NewLogReport(config)
}

func (df *DefaultInizializer) InitializePath(paths string) (*pathutils.PathResult, error) {
	return pathutils.GetPath(paths)
}
