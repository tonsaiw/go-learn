package main

import (
	"context"
	"fmt"

	"github.com/tonsaiw/go-learn/application"
)


func main() {
	app := application.New()
	err := app.Start(context.Background())
	if err != nil {
		fmt.Printf("Error starting application: %v\n", err)
	}
}
