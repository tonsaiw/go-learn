package application

import (
	"context"
	"fmt"
	"net/http"
)

// สำหรับจัดการ router และ server
type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		// จัดการ router
		router: loadRoutes(),

	}
	
	return app
}

func (a *App) Start(ctx context.Context) error {
	// จัดการ server
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}