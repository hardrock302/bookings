package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app_config *config.AppConfig) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(SessionLoad)
	router.Get("/", handlers.Repo.Home)
	router.Get("/home", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	return router
}
