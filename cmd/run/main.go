package main

import (
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

func main() {
	logger, err := infrastructure.InitLogger()
	if err != nil {
		fmt.Println("failed to initialize logger:", err)
		return
	}

	slog.SetDefault(logger.Logger)

	defer func() {
		if err := infrastructure.CloseLogger(logger); err != nil {
			fmt.Println("failed to close logger:", err)
		}
	}()

	initializer := &application.DefaultInizializer{}

	appComponents, err := application.InitializeApp(initializer)
	if err != nil {
		fmt.Println(err)
		return
	}

	dataProcessor, err := application.GetDataProcessor(appComponents.PathResult)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = dataProcessor.Process(appComponents.PathResult, &appComponents.LogReport)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputRenderer, err := application.GetOutputRenderer(appComponents.Config.Format)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputRenderer.Render(&appComponents.LogReport)
}
