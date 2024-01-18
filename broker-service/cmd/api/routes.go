package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Ling"},
		AllowCredentials: true,
		MaxAge:           3000,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", app.Broker)

	return mux
}
