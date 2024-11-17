package application

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/pathutils"
)

type AppComponents struct {
	Config     *domain.FlagConfig
	PathResult *pathutils.PathResult
	LogReport  domain.LogReport
}

func InitializeApp(initializer Inizializer) (*AppComponents, error) {
	config, err := initializer.InitializeConfig()
	if err != nil {
		return nil, fmt.Errorf("initializing config: %w", err)
	}

	pathResult, err := initializer.InitializePath(config.Path)
	if err != nil {
		return nil, fmt.Errorf("initializing path: %w", err)
	}

	logReport := initializer.InitializeLogReport(config)

	return &AppComponents{
		Config:     config,
		PathResult: pathResult,
		LogReport:  logReport,
	}, nil
}
