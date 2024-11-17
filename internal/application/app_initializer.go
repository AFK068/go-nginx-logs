package application

import (
	"fmt"
	"log/slog"

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
		slog.Error("failed to initialize config", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing config: %w", err)
	}

	pathResult, err := initializer.InitializePath(config.Path)
	if err != nil {
		slog.Error("failed to initialize path", slog.String("error", err.Error()))

		return nil, fmt.Errorf("initializing path: %w", err)
	}

	logReport := initializer.InitializeLogReport(config)

	slog.Info("app components initialized successfully, config:",
		slog.Any("config", config),
		slog.Any("pathResult", pathResult),
		slog.Any("logReport", logReport),
	)

	return &AppComponents{
		Config:     config,
		PathResult: pathResult,
		LogReport:  logReport,
	}, nil
}
