package main

import (
	"context"
	"fmt"

	"github.com/bagussans/ms-support-golang/application"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Application failed to start: %v\n", err)
	}
}
