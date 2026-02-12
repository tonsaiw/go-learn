package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tonsaiw/go-learn/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// จัดการ routes สำหรับ orders
	router.Route("/orders", loadOrderRoutes)

	return router
}

// จัดการ sub-routes สำหรับ orders
func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}
	router.Get("/", orderHandler.List)
	router.Post("/", orderHandler.Create)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.Update)
	router.Delete("/{id}", orderHandler.Delete)
}