package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	initializerMock "github.com/es-debug/backend-academy-2024-go-template/internal/application/mocks"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/pathutils"
	"github.com/stretchr/testify/assert"
)

func TestInitializeApp(t *testing.T) {
	mockInitializer := initializerMock.NewInizializer(t)

	config := &domain.FlagConfig{
		Path:   "some/path",
		Format: "markdown",
	}

	pathResult := &pathutils.PathResult{
		Paths: []string{"some/path"},
		Type:  "file",
	}

	logReport := domain.LogReport{
		FilterConfig: config,
	}

	mockInitializer.On("InitializeConfig").Return(config, nil)
	mockInitializer.On("InitializePath", config.Path).Return(pathResult, nil)
	mockInitializer.On("InitializeLogReport", config).Return(logReport)

	appComponents, err := application.InitializeApp(mockInitializer)

	assert.NoError(t, err)
	assert.NotNil(t, appComponents)
	assert.Equal(t, config, appComponents.Config)
	assert.Equal(t, pathResult, appComponents.PathResult)
	assert.Equal(t, logReport, appComponents.LogReport)

	mockInitializer.AssertExpectations(t)
}
