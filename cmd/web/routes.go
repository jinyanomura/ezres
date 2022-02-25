package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jinyanomura/ezres-web/pkg/config"
	"github.com/jinyanomura/ezres-web/pkg/handlers"
)

func routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/document", handlers.Repo.Document)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/example", handlers.Repo.Example)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}