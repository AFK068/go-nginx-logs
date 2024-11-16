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

func InitializeApp() (*AppComponents, error) {
	config, err := InitializeConfig()
	if err != nil {
		return nil, fmt.Errorf("initializing config: %w", err)
	}

	pathResult, err := InitializePath(config.Path)
	if err != nil {
		return nil, fmt.Errorf("initializing path: %w", err)
	}

	logReport := InitializeLogReport(config)

	return &AppComponents{
		Config:     config,
		PathResult: pathResult,
		LogReport:  logReport,
	}, nil
}
