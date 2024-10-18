package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/rolandtannous/gobookings/pkg/config"
	"github.com/rolandtannous/gobookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(noSurf)
	mux.Use(LoadAndSave)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return mux
}
