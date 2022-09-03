package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://localhost", "http://localhost"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Aceept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            false,
	}))

	// mux.Use(middleware.Heartbeat("/ping"))
	mux.Post("/all", app.Broker)
	mux.Post("/", app.Broker)

	mux.Post("/handle", app.HandleSubmission)
	mux.Post("/log-grpc", app.logViaGRPC)

	return mux

}
