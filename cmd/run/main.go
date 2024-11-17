package main

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
)

func main() {
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
